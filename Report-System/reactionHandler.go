{{/*
    This handy-dandy custom command-bundle allows a user to cancel their most recent report and utilizes
    Reactions to make things easier for staff.
    This custom command manages the reaction menu.
    Make this in a seperate Reaction CC, due to its massive character count.
    Remove this leading comment once you added this command.
    Obligatory Trigger type and trigger: Reaction; added reactions only.
    Created by: Olde#7325
*/}}
{{/*ACTUAL CODE*/}}
{{/*Initializing variables*/}}
{{$s := sdict (dbGet 2000 "reportSettings").Value}}
{{$rD := $s.reportDiscussion}}
{{$rL := $s.reportLog}}
{{/*Validation Steps*/}}
{{if eq .Channel.ID (toInt $rL)}}
{{$mod := (printf "\nResponsible moderator: <@%d>" .Reaction.UserID)}}
{{$isMod := in (split (index (split (exec "viewperms") "\n") 2) ", ") "ManageMessages"}}
{{if .ReactionMessage.Embeds}}
{{$e := (index .ReactionMessage.Embeds 0)}}
{{if and $e.Author $e.Footer}}
{{$r := index (getMessage $.Reaction.ChannelID $.Reaction.MessageID).Embeds 0|structToSdict}}
{{range $k, $v := $r}}{{if eq (kindOf $v true) "struct"}}{{$r.Set $k (structToSdict $v)}}{{end}}{{end}}
{{with $r}}
{{$e = (index $.ReactionMessage.Embeds 0)}}
{{if $isMod}}
{{if (reFind (toString $.User.ID) $e.Footer.Text)}}
{{$user := index (reFindAllSubmatches `\A<@!?(\d{17,19})>` .Description) 0 1|toInt|userArg}} {{/*Parsing user from description, saving a db call*/}}
{{.Set "Footer" (sdict "text" (print "Responsible Moderator: " $.User.String "(ID: " $.User.ID ")") "icon_url" ($.User.AvatarURL "256"))}}
{{.Set "Author" (sdict "name" (printf "%s: (ID %d)" $user.String $user.ID) "icon_url" ($user.AvatarURL "256"))}}
{{if eq $.Reaction.Emoji.Name "❌"}}{{/*Dismissal*/}}
{{sendMessage (toInt64 $rD) (printf "<@%d>: Your report was dismissed. %s" $user.ID $mod)}}
{{deleteAllMessageReactions nil $.Reaction.MessageID}}
{{.Set "Fields" ((cslice).AppendSlice $r.Fields)}}{{$r.Fields.Set 0 (sdict "name" "Current State" "value" "__Report dismissed.__")}}
{{.Set "Fields" ((cslice).AppendSlice $r.Fields)}}{{$r.Fields.Set 5 (sdict "name" "Reaction Menu Options" "value" "Warn for `False report` with ❗ or finish without warning with 👌.")}}
{{.Set "color" 65280}}
{{editMessage nil $.Reaction.MessageID (complexMessageEdit "embed" $r)}}
{{addReactions "❗" "👌"}}
{{dbSet $user.ID "key" "used"}}
{{else if eq $.Reaction.Emoji.Name "🛡️"}}{{/*Taking care*/}}
{{sendMessage (toInt64 $rD) (printf "<@%d>: Your report is being taken care of; Should you have any further information, please post it down below. %s" $user.ID $mod)}}
{{deleteAllMessageReactions nil $.Reaction.MessageID}}
{{.Set "Fields" ((cslice).AppendSlice $r.Fields)}}{{$r.Fields.Set 0 (sdict "name" "Current State" "value" "__Under investigation.__")}}
{{.Set "Fields" ((cslice).AppendSlice $r.Fields)}}{{$r.Fields.Set 5 (sdict "name" "Reaction Menu Options" "value" "Dismiss with ❌ or resolve with 👍.")}}
{{.Set "color" 16776960}}
{{editMessage nil $.Reaction.MessageID (complexMessageEdit "embed" $r)}}
{{addReactions "❌" "👍"}}
{{dbSet $user.ID "key" "used"}}
{{else if eq $.Reaction.Emoji.Name "⚠️"}}{{/*Request info*/}}
{{if ne (dbGet $user.ID "key").Value "used"}}{{/*Without cancellation request*/}}
{{sendMessage (toInt64 $rD) (printf "<@%d>: More information was requested. Please post it down below. %s" $user.ID $mod)}}
{{deleteAllMessageReactions nil $.Reaction.MessageID}}
{{.Set "Fields" ((cslice).AppendSlice $r.Fields)}}{{$r.Fields.Set 0 (sdict "name" "Current State" "value" "__More information requested.__")}}
{{.Set "Fields" ((cslice).AppendSlice $r.Fields)}}{{$r.Fields.Set 5 (sdict "name" "Reaction Menu Options" "value" "Dismiss with ❌ or start investigation with 🛡️.")}}
{{.Set "color" 255}}
{{editMessage nil $.Reaction.MessageID (complexMessageEdit "embed" $r)}}
{{addReactions "❌" "🛡️"}}
{{else}} 
{{/*With Cancellation request*/}}
{{sendMessage (toInt64 $rD) (printf "<@%d>: More information regarding your cancellation was requested. Please post it down below. %s" $user.ID $mod)}}
{{deleteAllMessageReactions nil $.Reaction.MessageID}}
{{.Set "Fields" ((cslice).AppendSlice $r.Fields)}}{{$r.Fields.Set 0 (sdict "name" "Current State" "value" "__More information requested.__")}}
{{.Set "Fields" ((cslice).AppendSlice $r.Fields)}}{{$r.Fields.Set 5 (sdict "name" "Reaction Menu Options" "value" "Dismiss request with 🚫, or accept request __(and nullify report)__ with ✅")}}
{{.Set "color" 255}}
{{editMessage nil $.Reaction.MessageID (complexMessageEdit "embed" $r)}}
{{addReactions "🚫" "✅"}}
{{end}}
{{else if eq $.Reaction.Emoji.Name "🚫"}}{{/*Dismissal of cancellation*/}}
{{sendMessage (toInt64 $rD) (printf "<@%d>: Your cancellation request was denied. %s" $user.ID $mod)}}
{{deleteAllMessageReactions nil $.Reaction.MessageID}}
{{.Set "Fields" ((cslice).AppendSlice $r.Fields)}}{{$r.Fields.Set 0 (sdict "name" "Current State" "value" "__Cancellation request denied.__")}}
{{.Set "Fields" ((cslice).AppendSlice $r.Fields)}}{{$r.Fields.Set 5 (sdict "name" "Reaction Menu Options" "value" "Dismiss report with ❌, put under investigation with 🛡️, or request more background information with ⚠️.")}}
{{.Set "color" 16711680}}
{{editMessage nil $.Reaction.MessageID (complexMessageEdit "embed" $r)}}
{{addReactions "❌" "🛡️" "⚠️"}}
{{else if eq $.Reaction.Emoji.Name "✅"}}{{/*Cancellation approved*/}}
{{sendMessage (toInt64 $rD) (printf "<@%d>: Your cancellation request was accepted. %s" $user.ID $mod)}}
{{deleteAllMessageReactions nil $.Reaction.MessageID}}
{{.Set "Fields" ((cslice).AppendSlice $r.Fields)}}{{$r.Fields.Set 0 (sdict "name" "Current State" "value" "__Cancellation request accepted, report nullified.__")}}
{{.Set "Fields" ((cslice).AppendSlice (slice $r.Fields 0 5))}}
{{.Set "Footer" (sdict "text" (print "Report closed! • Responsible Moderator: " $.User.String "(ID: " $.User.ID ")") "icon_url" ($.User.AvatarURL "256"))}}
{{.Set "color" 65280}}
{{editMessage nil $.Reaction.MessageID (complexMessageEdit "embed" $r)}}
{{addReactions "🏳️"}}
{{else if eq $.Reaction.Emoji.Name "👍"}}{{/*Report resolved*/}}
{{sendMessage (toInt64 $rD) (printf "<@%d>: Your report was resolved. %s" $user.ID $mod)}}
{{deleteAllMessageReactions nil $.Reaction.MessageID}}
{{.Set "Fields" ((cslice).AppendSlice $r.Fields)}}{{$r.Fields.Set 0 (sdict "name" "Current State" "value" "__Report resolved.__")}}
{{.Set "Fields" ((cslice).AppendSlice (slice $r.Fields 0 5))}}
{{.Set "Footer" (sdict "text" (print "Report closed! • Responsible Moderator: " $.User.String "(ID: " $.User.ID ")") "icon_url" ($.User.AvatarURL "256"))}}
{{.Set "color" 65280}}
{{editMessage nil $.Reaction.MessageID (complexMessageEdit "embed" $r)}}
{{addReactions "🏳️"}}
{{else if eq $.Reaction.Emoji.Name "❗"}}
{{$silent := exec "warn" $user.ID "False Report."}}
{{deleteAllMessageReactions nil $.Reaction.MessageID}}
{{.Set "Fields" ((cslice).AppendSlice $r.Fields)}}{{$r.Fields.Set 0 (sdict "name" "Current State" "value" "__Report dismissed, warned for false report.__")}}
{{.Set "Fields" ((cslice).AppendSlice (slice $r.Fields 0 5))}}
{{.Set "Footer" (sdict "text" (print "Report closed! • Responsible Moderator: " $.User.String "(ID: " $.User.ID ")") "icon_url" ($.User.AvatarURL "256"))}}
{{.Set "color" 65280}}
{{editMessage nil $.Reaction.MessageID (complexMessageEdit "embed" $r)}}
{{addReactions "🏳️"}}
{{else if eq $.Reaction.Emoji.Name "👌"}}
{{deleteAllMessageReactions nil $.Reaction.MessageID}}
{{.Set "Fields" ((cslice).AppendSlice $r.Fields)}}{{$r.Fields.Set 0 (sdict "name" "Current State" "value" "__Report dismissed, no further action taken.__")}}
{{.Set "Fields" ((cslice).AppendSlice (slice $r.Fields 0 5))}}
{{.Set "Footer" (sdict "text" (print "Report closed! • Responsible Moderator: " $.User.String "(ID: " $.User.ID ")") "icon_url" ($.User.AvatarURL "256"))}}
{{.Set "color" 65280}}
{{editMessage nil $.Reaction.MessageID (complexMessageEdit "embed" $r)}}
{{addReactions "🏳️"}}
{{end}}
{{else}}
{{deleteMessageReaction nil $.Reaction.MessageID $.User.ID "❌" "❗" "👌" "👍" "✅" "🛡️" "⚠️" "🚫"}}
{{if and (ne $.Reaction.Emoji.Name "🏳️") (reFind "•" $e.Footer.Text)}}
{{$tempMessage := sendMessageRetID nil (printf "<@%d>: No moderator yet, you claimed this report now. Your reactions were reset, please redo. Thanks ;)" $.User.ID)}}
{{deleteMessage nil $tempMessage 5}}
{{.Set "Footer" (sdict "text" (print "Responsible Moderator: " $.User.String " (ID: " $.User.ID ")") "icon_url" ($.User.AvatarURL "256"))}}
{{editMessage nil $.Reaction.MessageID (complexMessageEdit "embed" $r)}}{{end}}
{{end}}
{{else}}
{{deleteMessageReaction nil $.Reaction.MessageID $.User.ID "❌" "❗" "👌" "👍" "✅" "🛡️" "⚠️" "🚫"}}
{{end}}{{end}}{{end}}{{else}}{{end}}{{end}}
