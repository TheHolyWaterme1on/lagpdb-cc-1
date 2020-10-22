{{/*
    This handy-dandy custom command-bundle allows a user to cancel their most recent report and utilizes
    Reactions to make things easier for staff.
    This custom command manages and takes care of the cancellation requests.
    
    Recommended Trigger type and trigger: Regex; \A-c(ancel)?r(eport)?(\s+|\z)

    Credit: ye olde boi#7325 U-ID:665243449405997066
*/}}

{{/*CONFIG AREA START*/}}

{{$reports := 750730537571975298}} {{/*The channel where your reports are logged into.*/}}

{{/*CONFIG AREA END*/}}

{{/*ACTUAL CODE DO NOT TOUCH UNLESS YOU KNOW WHAT YOU ARE DOING*/}}
{{if not (ge (len .CmdArgs) 3)}}
    ```{{.Cmd}} <Message:ID> <Key:Text> <Reason:Text>```
    Not enough arguments passed.
{{else}}
    {{$dbValue := (dbGet .User.ID "key").Value|str}}
    {{$reportMessage := (index .CmdArgs 0)}}
    {{if (reFind (printf `\A<@!?%d>` .User.ID) (getMessage $reports $reportMessage).Content)}} 
            {{if eq "used" $dbValue}}
                Your latest report has already been cancelled!
            {{else}}
            {{if eq (index .CmdArgs 1|str) $dbValue}}
                {{if ge (len .CmdArgs) 3}}
                    {{$reason := joinStr " " (slice .CmdArgs 2)}}
                    {{$userReportString := (dbGet 2000 (print "userString-" .User.ID).Value)}}
                    {{$cancelGuide := (printf "\nDeny request with 🚫, accept with ✅, or request more information with ⚠️")}}
                    {{dbSet 2000 "cancelGuideBasic" $cancelGuide}}
                    {{$userCancelString := (printf "%s \n<@%d> requested cancellation of this report due to: `%s`" .User.ID $reason)}}
                    {{dbSet 2000 (print "userCancel-" .User.ID) $userCancelString}}
                    {{editMessage $reports $reportMessage (printf "%s %s. %s" $userReportString $userCancelString $cancelGuide)}}
                    Cancellation requested.
                    {{deleteAllMessageReactions $reports $reportMessage}}
                    {{addMessageReactions $reports $reportMessage "🚫" "✅" "⚠️"}}
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