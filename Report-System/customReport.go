{{/*
    This handy-dandy custom command-bundle allows a user to cancel their most recent report and utilizes
    Reactions to make things easier for staff.
    This custom command is basically the native report-command, but adds some back-end functionalites in order for the rest to work :)

    Usage: `-ru <User:Mention/ID> <Reason:Text>`

    Recommended Trigger type and trigger: Regex trigger with trigger `\A-r(eport)?u(ser)?(\s+|\z)`

    Created by: Olde#7325
*/}}

{{/*CONFIG AREA START*/}}

{{$reportLog := 772251753173221386}} {{/*The channel where your reports are logged into.*/}}
{{$reportDiscussion := 766370841196888104}} {{/*Your channel where users talk to staff*/}}

{{/*CONFIG AREA END*/}}


{{/*ACTUAL CODE*/}}
{{if (eq (len .CmdArgs) 1)}}
    {{if eq (index .CmdArgs 0) "dbSetup"}}
        {{if not ((dbGet 2000 "reportLog") (dbGet 2000 "reportDiscussion"))}}
            {{dbSet 2000 "reportLog" (toString $reportLog)}}
            {{dbSet 2000 "reportDiscussion" (toString $reportDiscussion)}}
            {{sendMessage nil "**Database primed, system is ready to use!**"}}
        {{else}}
            {{sendMessage nil "**Database entries already exist! No action taken, system still ready to use.**"}}
        {{end}}
    {{end}}
{{else if not (ge (len .CmdArgs) 2)}}
    ```{{.Cmd}} <User:Mention/ID> <Reason:Text>```
    Not enough arguments passed.
{{else}}
    {{$secret := adjective}}
    {{$s := execAdmin "log"}}
    {{$user := userArg (index .CmdArgs 0)}}
    {{$reason := joinStr " " (slice .CmdArgs 1)}}
    {{$reportGuide := (printf "\nDismiss report with ❌, take action with 🛡️, or request more background information with ⚠️")}}
    {{$userReportString := (printf  "<@%d> reported <@%d> in <#%d> for: `%s` \n Last 100 messages: <%s>" .User.ID $user.ID .Channel.ID $reason $s)}}
    {{dbSet 2000 "reportGuideBasic" $reportGuide}}
    {{dbSet 2000 (printf "userReport%d" .User.ID) $userReportString}}
    {{$x := sendMessageRetID $reportLog (printf "%s %s" $userReportString $reportGuide)}}
    {{addMessageReactions $reportLog $x "❌" "🛡️" "⚠️"}}
    User reported to the proper authorites!
    {{dbSet .User.ID "key" $secret}}
    {{sendDM (printf "User reported to the proper authorities! If you wish to cancel your report, simply type `-cancelr %d %s` in any channel.\n **A reason is required.**" $x $secret)}}
{{end}}