{{/*
    This handy-dandy custom command-bundle allows a user to cancel their most recent report and utilizes
    Reactions to make things easier for staff.
    This custom command manages and takes care of the cancellation requests.

    Usage: `-cr <Message:ID> <Key:Text> <Reason:Text>`
    
    Recommended Trigger type and trigger: Regex trigger with trigger `\A-c(ancel)?r(eport)?(\s+|\z)`

    Created by: Olde#7325
*/}}


{{/*CONFIG AREA START*/}}

{{$reportLog := 750730537571975298}} {{/*The channel where your reports are logged into.*/}}
{{$reportDiscussion := 750099460314628176}} {{/*Your channel where users talk to staff*/}}

{{/*CONFIG AREA END*/}}


{{/*ACTUAL COMMENT*/}}
{{if not (ge (len .CmdArgs) 3)}}
    ```{{.Cmd}} <Message:ID> <Key:Text> <Reason:Text>```
    Not enough arguments passed.
{{else}}
    {{$dbValue := (dbGet .User.ID "key").Value|str}}
    {{$reportMessage := ((index .CmdArgs 0)|toInt64)}}
    {{$reportMessageContent := (getMessage $reportLog $reportMessage).Content}}
    {{if (reFind (printf `\A<@!?%d>` .User.ID) $reportMessageContent)}} 
            {{if eq "used" $dbValue}}
                Your latest report has already been cancelled!
            {{else}}
            {{if eq (index .CmdArgs 1|str) $dbValue}}
                {{if ge (len .CmdArgs) 3}}
                    {{$reason := joinStr " " (slice .CmdArgs 2)}}
                    {{$userReportString := (dbGet 2000 (printf "userReport%d" .User.ID)).Value|str}}
                    {{$cancelGuide := (printf "Deny request with 🚫, accept with ✅, or request more information with ⚠️")}}
                    {{dbSet 2000 "cancelGuideBasic" $cancelGuide}}
                    {{$userCancelString := (printf "<@%d> requested cancellation of this report due to: `%s`" .User.ID $reason)}}
                    {{dbSet 2000 (printf "userCancel%d" .User.ID) $userCancelString}}
                    {{editMessage $reportLog $reportMessage (printf "%s \n %s. \n %s" $userReportString $userCancelString $cancelGuide)}}
                    Cancellation requested, have a nice day!
                    {{deleteAllMessageReactions $reportLog $reportMessage}}
                    {{addMessageReactions $reportLog $reportMessage "🚫" "✅" "⚠️"}}
                    {{dbSet .User.ID "key" "used"}}
                {{end}}
            {{else}}
                Invalid key provided!
            {{end}}
        {{end}}
        {{else}}
            You are not the author of this report!
    {{end}}
{{end}}