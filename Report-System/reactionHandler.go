{{/*
    This handy-dandy custom command-bundle allows a user to cancel their most recent report and utilizes
    Reactions to make things easier for staff.
    This custom command manages the reaction menu.
    You can put this code into your reaction listener, should you already have one. If not, make a new one ;)
    Remove this leading comment once you added this command to save on character count.

    Recommended Trigger type and trigger: Reaction; added and removed reactions.

    Created by: Olde#7325
*/}}

{{/*CONFIG AREA START*/}}

{{$reportLog := 750730537571975298}} {{/*The channel where your reports are logged into.*/}}
{{$reportDiscussion := 750099460314628176}} {{/*Your channel where users talk to staff*/}}

{{/*CONFIG AREA END*/}}

{{/*ACTUAL CODE*/}}

{{/*Validation steps*/}}
{{if .Reaction}}
{{if .ReactionAdded}}
{{if eq .Reaction.ChannelID $reportLog}}

{{/*Set some vars, cutting down on DB stuff, Readability shit*/}}

{{$reportGuide := ((dbGet 2000 "reportGuideBasic").Value|str)}}
{{$user := (index (reFindAllSubmatches `\A(?:<@!?)?(\d{17,19})(?:>)?` .ReactionMessage.Content) 0 1|toInt64)}}
{{$userReportString := ((dbGet 2000 (printf "userReport%d" $user)).Value|str)}}
{{$userCancelString := ((dbGet 2000 (printf "userCancel%d" $user)).Value|str)}}
{{$mod := (printf "\nResponsible moderator: <@%d>" .Reaction.UserID)}}

{{if dbGet .Reaction.MessageID "ModeratorID"}}
    {{if eq .User.ID ((dbGet .Reaction.MessageID "ModeratorID").Value|toInt64)}}
        {{if eq .Reaction.Emoji.Name "❌"}}{{/*Dismissal*/}}
            {{sendMessage $reportDiscussion (printf "<@%d>: Your report has been dismissed. %s" $user $mod)}}
            {{deleteAllMessageReactions nil .Reaction.MessageID}}
            {{editMessage $reportLog .Reaction.MessageID (printf "%s\n **Report dismissed.** %s \nWarn for `false report` with ❗ or finish without warning with 👌." $userReportString $mod)}}
            {{addReactions "❗" "👌"}}
            {{dbSet $user "key" "used"}}
        {{else if eq .Reaction.Emoji.Name "🛡️"}}{{/*Taking care*/}}
            {{sendMessage $reportDiscussion (printf "<@%d>: Your report is being taken care of; Should you have any further information, please post it down below. %s" $user $mod)}}
            {{deleteAllMessageReactions nil .Reaction.MessageID}}
            {{editMessage $reportLog .Reaction.MessageID (printf "%s\n **Under investigation.** %s \nDismiss with ❌ or resolve with 👍." $userReportString $mod)}}
            {{addReactions "❌" "👍"}}
            {{dbSet $user "key" "used"}}
        {{else if eq .Reaction.Emoji.Name "⚠️"}}{{/*Request info*/}}
            {{if not (eq ((dbGet $user "key").Value) "used")}}{{/*Without cancellation request*/}}
                {{sendMessage $reportDiscussion (printf "<@%d>: More information has been requested. Please post it down below. %s" $user $mod)}}
                {{deleteAllMessageReactions nil .Reaction.MessageID}}
                {{editMessage $reportLog .Reaction.MessageID (printf "%s\n **More information requested.** %s \nDismiss with ❌ or start investigation with 🛡️." $userReportString $mod)}}
                {{addReactions "❌" "🛡️"}}
            {{else}} 
            {{/*With Cancellation request*/}}
                {{sendMessage $reportDiscussion (printf "<@%d>: More information regarding your cancellation has been requested. Please post it down below. %s" $user $mod)}}
                {{deleteAllMessageReactions nil .Reaction.MessageID}}
                {{editMessage $reportLog .Reaction.MessageID (printf "%s\n \n%s \n**More information requested.** %s \nDismiss request with 🚫, or accept request __(and nullify report)__ with ✅" $userReportString $userCancelString $mod)}}
                {{addReactions "🚫" "✅"}}
            {{end}}
        {{else if eq .Reaction.Emoji.Name "🚫"}}{{/*Dismissal of cancellation*/}}
            {{sendMessage $reportDiscussion (printf "<@%d>: Your request of cancellation has been dismissed. %s" $user $mod)}}
            {{deleteAllMessageReactions nil .Reaction.MessageID}}
            {{editMessage $reportLog .Reaction.MessageID (printf "%s\n %s\n**Cancellation request denied.** %s \n%s" $userReportString $userCancelString $mod $reportGuide)}}
            {{addReactions "❌" "🛡️" "⚠️"}}
        {{else if eq .Reaction.Emoji.Name "✅"}}{{/*Cancellation approved*/}}
            {{sendMessage $reportDiscussion (printf "<@%d>: Your request of cancellation has been accepted. %s" $user $mod)}}
            {{deleteAllMessageReactions nil .Reaction.MessageID}}
            {{editMessage $reportLog .Reaction.MessageID (printf "%s\n %s **Cancellation request accepted. Report nullified.** %s" $userReportString $userCancelString $mod)}}
            {{addReactions "🏳️"}}
        {{else if eq .Reaction.Emoji.Name "👍"}}{{/*Report resolved*/}}
            {{sendMessage $reportDiscussion (printf "<@%d>: Your report has been resolved. %s" $user $mod)}}
            {{deleteAllMessageReactions nil .Reaction.MessageID}}
            {{editMessage $reportLog .Reaction.MessageID (printf "%s\n **Report resolved.** %s" $userReportString $mod)}}
            {{addReactions "🏳️"}}
        {{else if eq .Reaction.Emoji.Name "❗"}}
            {{$silent := exec "warn" $user "False Report."}}
            {{deleteAllMessageReactions nil .Reaction.MessageID}}
            {{editMessage $reportLog .Reaction.MessageID (printf "%s\n **Report dismissed. Warned for False report.** %s" $userReportString $mod)}}
            {{addReactions "🏳️"}}
        {{else if eq .Reaction.Emoji.Name "👌"}}
            {{deleteAllMessageReactions nil .Reaction.MessageID}}
            {{editMessage $reportLog .Reaction.MessageID (printf "%s \n **Report dismissed. No action taken.** %s" $userReportString $mod)}}
        {{else if eq .Reaction.Emoji.Name "🏳️"}}
            {{deleteMessageReaction nil .Reaction.MessageID .User.ID "🏳️"}}
        {{end}}
    {{else}}
        {{deleteMessageReaction nil .Reaction.MessageID .User.ID "❌" "❗" "👌" "👍" "✅" "🛡️" "⚠️" "🚫"}}
    {{end}}
    {{dbSet .Reaction.MessageID "ModeratorID" .User.ID}}
    {{deleteMessageReaction nil .Reaction.MessageID .User.ID "❌" "❗" "👌" "👍" "✅" "🛡️" "⚠️" "🚫"}}
    {{$tempMessage := sendMessageRetID nil (printf "<@%d>: No moderator detected, you claimed this report now. Your reactions were reset, please redo. Thanks ;)" .User.ID)}}
    {{deleteMessage nil $tempMessage 15}}
{{end}}
{{end}}{{end}}{{end}}