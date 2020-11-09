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
{{/*Validation steps*/}}
{{$reportLog := (dbGet 2000 "reportLog").Value|toInt64}}
{{$reportDiscussion := (dbGet 2000 "reportDiscussion").Value|toInt64}}
{{if eq .Reaction.ChannelID $reportLog}}
{{/*Set some vars, cutting down on DB stuff, Readability shit*/}}
{{$reportGuide := ((dbGet 2000 "reportGuideBasic").Value|str)}}{{$user := (toInt64 (dbGet .Reaction.MessageID "reportAuthor").Value)}}{{$userReportString := ((dbGet 2000 (printf "userReport%d" $user)).Value|str)}}
{{$userCancelString := ((dbGet 2000 (printf "userCancel%d" $user)).Value|str)}}{{$mod := (printf "\nResponsible moderator: <@%d>" .Reaction.UserID)}}{{$modRoles := (cslice).AppendSlice (dbGet 2000 "modRoles").Value}}
{{$isMod := false}} {{range .Member.Roles}} {{if in $modRoles .}} {{$isMod = true}}{{end}}{{end}}
{{$report := index (getMessage nil .Reaction.MessageID).Embeds 0|structToSdict}}{{range $k, $v := $report}}{{if eq (kindOf $v true) "struct"}}{{$report.Set $k (structToSdict $v)}}{{end}}{{end}}
{{$report.Author.Set "icon_url" .Author.IconURL}}

{{if $isMod}}
    {{$report.Set "Footer" (sdict "text" (print "Responsible Moderator: " .User.String) "icon_url" (.User.AvatarURL "256"))}}
    {{if (dbGet .Reaction.MessageID "ModeratorID")}}
        {{if eq .User.ID (toInt64 (dbGet .Reaction.MessageID "ModeratorID").Value)}}
            {{if eq .Reaction.Emoji.Name "❌"}}{{/*Dismissal*/}}
                {{sendMessage $reportDiscussion (printf "<@%d>: Your report was dismissed. %s" $user $mod)}}
                {{deleteAllMessageReactions nil .Reaction.MessageID}}
                {{$report.Set "Fields" ((cslice).AppendSlice $report.Fields)}}{{$report.Fields.Set 0 (sdict "name" "Current State" "value" "__Report dismissed.__")}}
                {{$report.Set "Fields" ((cslice).AppendSlice $report.Fields)}}{{$report.Fields.Set 4 (sdict "name" "Reaction Menu Options" "value" "Warn for `False report` with ❗ or finish without warning with 👌.")}}
                {{$report.Set "color" 65280}}
                {{editMessage nil .Reaction.MessageID (complexMessageEdit "embed" $report)}}
                {{addReactions "❗" "👌"}}
                {{dbSet $user "key" "used"}}
            {{else if eq .Reaction.Emoji.Name "🛡️"}}{{/*Taking care*/}}
                {{sendMessage $reportDiscussion (printf "<@%d>: Your report is being taken care of; Should you have any further information, please post it down below. %s" $user $mod)}}
                {{deleteAllMessageReactions nil .Reaction.MessageID}}
                {{$report.Set "Fields" ((cslice).AppendSlice $report.Fields)}}{{$report.Fields.Set 0 (sdict "name" "Current State" "value" "__Under investigation.__")}}
                {{$report.Set "Fields" ((cslice).AppendSlice $report.Fields)}}{{$report.Fields.Set 4 (sdict "name" "Reaction Menu Options" "value" "Dismiss with ❌ or resolve with 👍.")}}
                {{$report.Set "color" 16776960}}
                {{editMessage nil .Reaction.MessageID (complexMessageEdit "embed" $report)}}
                {{addReactions "❌" "👍"}}
                {{dbSet $user "key" "used"}}
            {{else if eq .Reaction.Emoji.Name "⚠️"}}{{/*Request info*/}}
                {{if not (eq ((dbGet $user "key").Value) "used")}}{{/*Without cancellation request*/}}
                    {{sendMessage $reportDiscussion (printf "<@%d>: More information has been requested. Please post it down below. %s" $user $mod)}}
                    {{deleteAllMessageReactions nil .Reaction.MessageID}}
                    {{$report.Set "Fields" ((cslice).AppendSlice $report.Fields)}}{{$report.Fields.Set 0 (sdict "name" "Current State" "value" "__More information requested.__")}}
                    {{$report.Set "Fields" ((cslice).AppendSlice $report.Fields)}}{{$report.Fields.Set 4 (sdict "name" "Reaction Menu Options" "value" "Dismiss with ❌ or start investigation with 🛡️.")}}
                    {{$report.Set "color" 255}}
                    {{editMessage nil .Reaction.MessageID (complexMessageEdit "embed" $report)}}
                    {{addReactions "❌" "🛡️"}}
                {{else}} 
                    {{/*With Cancellation request*/}}
                    {{sendMessage $reportDiscussion (printf "<@%d>: More information regarding your cancellation has been requested. Please post it down below. %s" $user $mod)}}
                    {{deleteAllMessageReactions nil .Reaction.MessageID}}
                    {{$report.Set "Fields" ((cslice).AppendSlice $report.Fields)}}{{$report.Fields.Set 0 (sdict "name" "Current State" "value" "__More information requested.__")}}
                    {{$report.Set "Fields" ((cslice).AppendSlice $report.Fields)}}{{$report.Fields.Set 4 (sdict "name" "Reaction Menu Options" "value" "Dismiss request with 🚫, or accept request __(and nullify report)__ with ✅")}}
                    {{$report.Set "color" 255}}
                    {{editMessage nil .Reaction.MessageID (complexMessageEdit "embed" $report)}}
                    {{addReactions "🚫" "✅"}}
                {{end}}
            {{else if eq .Reaction.Emoji.Name "🚫"}}{{/*Dismissal of cancellation*/}}
                {{sendMessage $reportDiscussion (printf "<@%d>: Your request of cancellation has been dismissed. %s" $user $mod)}}
                {{deleteAllMessageReactions nil .Reaction.MessageID}}
                {{$report.Set "Fields" ((cslice).AppendSlice $report.Fields)}}{{$report.Fields.Set 0 (sdict "name" "Current State" "value" "__Cancellation request denied.__")}}
                {{$report.Set "Fields" ((cslice).AppendSlice $report.Fields)}}{{$report.Fields.Set 4 (sdict "name" "Reaction Menu Options" "value" $reportGuide)}}
                {{$report.Set "color" 16711680}}
                {{editMessage nil .Reaction.MessageID (complexMessageEdit "embed" $report)}}
                {{addReactions "❌" "🛡️" "⚠️"}}
            {{else if eq .Reaction.Emoji.Name "✅"}}{{/*Cancellation approved*/}}
                {{sendMessage $reportDiscussion (printf "<@%d>: Your request of cancellation has been accepted. %s" $user $mod)}}
                {{deleteAllMessageReactions nil .Reaction.MessageID}}
                {{$report.Set "Fields" ((cslice).AppendSlice $report.Fields)}}{{$report.Fields.Set 0 (sdict "name" "Current State" "value" "__Cancellation request accepted, report nullified.__")}}
                {{$report.Set "Fields" ((cslice).AppendSlice (slice $report.Fields 0 3))}}
                {{$report.Set "color" 65280}}
                {{editMessage nil .Reaction.MessageID (complexMessageEdit "embed" $report)}}
                {{dbDel .Reaction.MessageID "ModeratorID"}}
                {{addReactions "🏳️"}}
            {{else if eq .Reaction.Emoji.Name "👍"}}{{/*Report resolved*/}}
                {{sendMessage $reportDiscussion (printf "<@%d>: Your report has been resolved. %s" $user $mod)}}
                {{deleteAllMessageReactions nil .Reaction.MessageID}}
                {{$report.Set "Fields" ((cslice).AppendSlice $report.Fields)}}{{$report.Fields.Set 0 (sdict "name" "Current State" "value" "__Report resolved.__")}}
                {{$report.Set "Fields" ((cslice).AppendSlice (slice $report.Fields 0 3))}}
                {{$report.Set "color" 65280}}
                {{editMessage nil .Reaction.MessageID (complexMessageEdit "embed" $report)}}
                {{dbDel .Reaction.MessageID "ModeratorID"}}
                {{addReactions "🏳️"}}
            {{else if eq .Reaction.Emoji.Name "❗"}}
                {{$silent := exec "warn" $user "False Report."}}
                {{deleteAllMessageReactions nil .Reaction.MessageID}}
                {{$report.Set "Fields" ((cslice).AppendSlice $report.Fields)}}{{$report.Fields.Set 0 (sdict "name" "Current State" "value" "__Report dismissed, warned for false report.__")}}
                {{$report.Set "Fields" ((cslice).AppendSlice (slice $report.Fields 0 3))}}
                {{editMessage nil .Reaction.MessageID (complexMessageEdit "embed" $report)}}
                {{dbDel .Reaction.MessageID "ModeratorID"}}
                {{addReactions "🏳️"}}
            {{else if eq .Reaction.Emoji.Name "👌"}}
                {{deleteAllMessageReactions nil .Reaction.MessageID}}
                {{$report.Set "Fields" ((cslice).AppendSlice $report.Fields)}}{{$report.Fields.Set 0 (sdict "name" "Current State" "value" "__Report dismissed, no further action taken.__")}}
                {{$report.Set "Fields" ((cslice).AppendSlice (slice $report.Fields 0 3))}}
                {{editMessage nil .Reaction.MessageID (complexMessageEdit "embed" $report)}}
                {{dbDel .Reaction.MessageID "ModeratorID"}}
                {{addReactions "🏳️"}}
            {{end}}
        {{else}}
            {{deleteMessageReaction nil .Reaction.MessageID .User.ID "❌" "❗" "👌" "👍" "✅" "🛡️" "⚠️" "🚫"}}
        {{end}}
    {{else}}
        {{if ne .Reaction.Emoji.Name "🏳️"}}
        {{dbSet .Reaction.MessageID "ModeratorID" (toString .User.ID)}}
        {{deleteMessageReaction nil .Reaction.MessageID .User.ID "❌" "❗" "👌" "👍" "✅" "🛡️" "⚠️" "🚫"}}
        {{$tempMessage := sendMessageRetID nil (printf "<@%d>: No moderator detected, you claimed this report now. Your reactions were reset, please redo. Thanks ;)" .User.ID)}}
        {{deleteMessage nil $tempMessage 15}}
        {{$report.Set "Footer" (sdict "text" (print "Responsible Moderator: " .User.String) "icon_url" (.User.AvatarURL "256"))}}
        {{editMessage nil .Reaction.MessageID (complexMessageEdit "embed" $report)}}{{end}}
    {{end}}
{{else}}
{{deleteMessageReaction nil .Reaction.MessageID .User.ID "❌" "❗" "👌" "👍" "✅" "🛡️" "⚠️" "🚫"}}
{{end}}{{end}}