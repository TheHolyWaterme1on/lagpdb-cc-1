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
{{$modRoles := cslice 766372666483408947 766372758799122442}} {{/*RoleIDs of the roles which are considered moderator.*/}}
{{$adminRoles := cslice 766372666483408947}} {{/*RoleIDs of the roles which are considered admins. Can prime the database to setup the system and reset report count.*/}}

{{/*CONFIG AREA END*/}}


{{/*ACTUAL CODE*/}}
{{$isAdmin := false}} {{range .Member.Roles}} {{if in $adminRoles .}} {{$isAdmin = true}} {{end}} {{end}}
{{if (eq (len .CmdArgs) 1)}}
    {{if eq (index .CmdArgs 0) "dbSetup"}}
        {{if $isAdmin}}
            {{if not (and (dbGet 2000 "reportLog") (dbGet 2000 "reportDiscussion") (dbGet 2000 "modRoles") (dbGet 2000 "adminRoles") (eq (toInt (dbGet 2000 "ReportNo").Value) 0))}}
                {{dbSet 2000 "reportLog" (toString $reportLog)}}
                {{dbSet 2000 "reportDiscussion" (toString $reportDiscussion)}}
                {{dbSet 2000 "modRoles" $modRoles}}
                {{dbSet 2000 "adminRoles" $adminRoles}}
                {{dbSet 2000 "ReportNo" 0}}
                {{sendMessage nil "**Database primed, report numbers resetted, system is ready to use!**"}}
            {{else}}
                {{sendMessage nil "**Database entries already exist! No action taken, system still ready to use.**"}}
            {{end}}
        {{else}}
            {{sendMessage nil "You do not have permission to use this command!"}}
        {{end}}
    {{end}}
{{else if not (ge (len .CmdArgs) 2)}}
    {{sendMessage nil (printf "```%s <User:Mention/ID> <Reason:Text>``` \n Not enough arguments passed." .Cmd)}}
{{else}}
    {{$user := userArg (index .CmdArgs 0)}}
    {{if eq $user.ID .User.ID}}
        {{$silly := sendMessageRetID nil "You can't report yourself, silly."}}
        {{deleteTrigger}}
        {{deleteMessage nil $silly}}
    {{else}}
        {{$secret := adjective}}
        {{$logs250 := execAdmin "log" "250"}}
        {{$reason := joinStr " " (slice .CmdArgs 1)}}
        {{$reportGuide := (printf "\nDismiss report with ❌, put under investigation with 🛡️, or request more background information with ⚠️.")}}
        {{$userReportString := (printf  "<@%d> reported <@%d> in <#%d>." .User.ID $user.ID .Channel.ID)}}
        {{dbSet 2000 "reportGuideBasic" $reportGuide}}
        {{dbSet 2000 (printf "userReport%d" .User.ID) $userReportString}}
        {{$reportNo := dbIncr 2000 "ReportNo" 1}}
        {{$reportEmbed := cembed "title" (print "Report No. " $reportNo)
            "author" (sdict "name" (printf "%s (ID %d)" .User.String .User.ID) "icon_url" (.User.AvatarURL "256"))
            "thumbnail" (sdict "url" ($user.AvatarURL "512"))
            "description" $userReportString
            "fields" (cslice
                (sdict "name" "Current State" "value" "__Not reviewed yet.__")
                (sdict "name" "Reason for Report" "value" $reason)
                (sdict "name" "Reported user" "value" (printf "<@%d> (ID %d)" $user.ID $user.ID))
                (sdict "name" "Message Logs" "value" (printf "[last 250 messages](%s) \nTime - `%s`" $logs250 (currentTime.Format "Mon 02 Jan 15:04:05")))
                (sdict "name" "Reaction Menu Options" "value" $reportGuide)
            )
            "footer" (sdict "text" "No moderator yet • Claim with any reaction")
        }}
        {{$x := sendMessageRetID $reportLog $reportEmbed}}
        {{addMessageReactions $reportLog $x "❌" "🛡️" "⚠️"}}
        {{sendMessage nil "User reported to the proper authorites!"}}
        {{dbSet .User.ID "key" $secret}}
        {{dbSet $x "reportAuthor" (toString .User.ID)}}
        {{deleteTrigger}}
        {{deleteResponse}}
        {{sendDM (printf "User reported to the proper authorities! If you wish to cancel your report, simply type `-cancelr %d %s` in any channel.\n **A reason is required.**" $x $secret)}}
    {{end}}
{{end}}