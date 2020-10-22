{{/*
    This handy-dandy custom command-bundle allows a user to cancel their most recent report and utilizes
    Reactions to make things easier for staff.
    This custom command manages the reaction menu.
    You can put this code into your reaction listener, should you already have one. If not, make a new one ;)
    Should you get an error because your reaction CC is too long, simply remove all of yours (and mine) comments.

    Recommended Trigger type and trigger: Reaction; added and removed reactions.

    Credit: ye olde boi#7325 U-ID:665243449405997066
*/}}

{{/*CONFIG AREA START*/}}

{{/* REPORT MESSAGE REACTION HANDLER */}}
{{$reports := 750730537571975298}} {{/*Your channel reports shall be logged to.*/}}
{{$reportLog := 750099460314628176}} {{/*Your channel where users talk to staff*/}}

{{/**CONFIG AREA END*/}}

{{/*ACTUAL CODE DO NOT TOUCH UNLESS YOU KNOW WHAT YOU DO*/}}
{{if .Reaction}}
{{if eq .ReactionAdded true}}
{{if eq .Channel.ID $reports}} {{/*Validation steps*/}}

{{$reportGuide := ((dbGet 2000 "reportGuideBasic").Value|str)}}
{{$user := (index (reFindAllSubmatches `\A(?:<@!?)?(\d{17,19})(?:>)?` .ReactionMessage.Content) 0 1|toInt64)}}
{{$userReportString := ((dbGet 2000 (print "userReport-" $user)).Value|str)}}
{{$userCancelString := ((dbGet 2000 (print "userCancel-" $user)).Value|str)}}
{{$mod := (printf "\nResponsible moderator: <@%d>" .Reaction.UserID)}} {{/*Set some vars, cutting down on DB stuff, Readability shit*/}}

{{if eq .Reaction.Emoji.Name "❌"}}{{/*Dismissal*/}}
{{sendMessage $reportLog (printf "<@%d>: Your report has been dismissed. %s" $user $mod)}}
{{deleteAllMessageReactions nil .Reaction.MessageID}}
{{editMessage $reports .Reaction.MessageID (printf "%s\n **Report dismissed.** %s \nWarn for `false report` with ❗ or finish without warning with 👌.") $userReportString $mod}}
{{addReactions "❗" "👌"}}
{{dbSet $user "key" "used"}}
{{else if eq .Reaction.Emoji.Name "🛡️"}}{{/*Taking care*/}}
{{sendMessage $reportLog (printf "<@%d>: Your report is being taken care of; Should you have any further information, please post it down below. %s" $user $mod)}}
{{deleteAllMessageReactions nil .Reaction.MessageID}}
{{editMessage $reports .Reaction.MessageID (printf "%s\n **Under investigation.** %s \nDismiss with ❌ or resolve with 👍." $userReportString $mod)}}
{{addReactions "❌" "👍"}}
{{dbSet $user "key" "used"}}
{{else if eq .Reaction.Emoji.Name "⚠️"}}{{/*Request info*/}}
{{if not (eq ((dbGet $user "key").Value) "used")}}{{/*No cancellation request*/}}
{{sendMessage $reportLog (printf "<@%d>: More information has been requested. Please post it down below. %s" $user $mod)}}
{{deleteAllMessageReactions nil .Reaction.MessageID}}
{{editMessage $reports .Reaction.MessageID (printf "%s\n **More information requested.** %s \nDismiss with ❌ or start investigation with 🛡️." $userReportString $mod)}}
{{addReactions "❌" "🛡️"}}
{{else}}{{/*Cancellation request*/}}
{{sendMessage $reportLog (printf "<@%d>: More information regarding your cancellation has been requested. Please post it down below. %s" $user $mod)}}
{{deleteAllMessageReactions nil .Reaction.MessageID}}
{{editMessage $reports .Reaction.MessageID (printf "%s\n \n%s \n**More information requested.** %s \nDismiss request with 🚫, or accept request __(and nullify report)__ with ✅" $userReportString $userCancelString $mod)}}
{{addReactions "🚫" "✅"}}
{{end}}
{{else if eq .Reaction.Emoji.Name "🚫"}}{{/*Dismissal of cancellation*/}}
{{sendMessage $reportLog (printf "<@%d>: Your request of cancellation has been dismissed. %s" $user $mod)}}
{{deleteAllMessageReactions nil .Reaction.MessageID}}
{{editMessage $reports .Reaction.MessageID (printf "%s\n %s\n**Cancellation request denied.** %s \n%s" $userReportString $userCancelString $mod $reportGuide)}}
{{addReactions "❌" "🛡️" "⚠️"}}
{{else if eq .Reaction.Emoji.Name "✅"}}{{/*Cancellation approved*/}}
{{sendMessage $reportLog (printf "<@%d>: Your request of cancellation has been accepted. %s" $user $mod)}}
{{deleteAllMessageReactions nil .Reaction.MessageID}}
{{editMessage $reports .Reaction.MessageID (printf "%s\n %s **Cancellation request accepted. Report nullified.** %s" $userReportString $userCancelString $mod)}}
{{addReactions "🏳️"}}
{{else if eq .Reaction.Emoji.Name "👍"}}{{/*Report resolved*/}}
{{sendMessage $reportLog (printf "<@%d>: Your report has been resolved. %s" $user $mod)}}
{{deleteAllMessageReactions nil .Reaction.MessageID}}
{{editMessage $reports .Reaction.MessageID (printf "%s\n **Report resolved.** %s" $userReportString $mod)}}
{{addReactions "🏳️"}}
{{else if eq .Reaction.Emoji.Name "❗"}}
{{exec "warn" $user "False Report."}}
{{deleteAllMessageReactions nil .Reaction.MessageID}}
{{editMessage $reports .Reaction.MessageID (printf "%s\n **Report dismissed. Warned for False report.** %s" $userReportString $mod)}}
{{addReactions "🏳️"}}
{{else if eq .Reaction.Emoji.Name "👌"}}
{{deleteAllMessageReactions}}
{{editMessage $reports .Reaction.MessageID (printf "%s \n **Report dismissed. No action taken.** 5s" $userReportString $mod)}}
{{else if eq .Reaction.Emoji.Name "🏳️"}}
{{deleteMessageReaction nil .Reaction.MessageID .User.ID "🏳️"}}
{{end}}{{end}}{{end}}{{end}}