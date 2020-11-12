{{/*
    This handy-dandy custom command-bundle allows a user to cancel their most recent report and utilizes
    Reactions to make things easier for staff.
    This custom command manages and takes care of the cancellation requests.

    Usage: `-cr <Message:ID> <Key:Text> <Reason:Text>`
    
    Recommended Trigger type and trigger: Regex trigger with trigger `\A-c(ancel)?r(eport)?(\s+|\z)`

    Created by: Olde#7325
*/}}

{{/*ACTUAL CODE*/}}
{{if not (ge (len .CmdArgs) 3)}}
    ```{{.Cmd}} <Message:ID> <Key:Text> <Reason:Text>```
    Not enough arguments passed.
{{else}}
    {{$reportLog := (dbGet 2000 "reportLog").Value|toInt64}}
    {{$reportDiscussion := (dbGet 2000 "reportDiscussion").Value|toInt64}}
    {{$userKey := (dbGet .User.ID "key").Value|str}}
    {{$reportMessageID := ((index .CmdArgs 0)|toInt64)}}
    {{if eq (toInt64 (dbGet $reportMessageID "reportAuthor").Value) (toInt64 .User.ID)}}
            {{if eq "used" $userKey}}
                Your latest report already has been cancelled!
            {{else}}
            {{if eq (index .CmdArgs 1|str) $userKey}}
                {{if ge (len .CmdArgs) 3}}
                    {{$reason := joinStr " " (slice .CmdArgs 2)}}
                    {{$userReportString := (dbGet 2000 (printf "userReport%d" .User.ID)).Value|str}}
                    {{$cancelGuide := (printf "Deny request with 🚫, accept with ✅, or request more information with ⚠️.")}}
                    {{dbSet 2000 "cancelGuideBasic" $cancelGuide}}
                    {{$userCancelString := (printf "Cancellation of this report was requested. \n Reason: `%s`" $reason)}}
                    {{$combinedString := (print $userReportString " \n " $userCancelString)}}
                    {{dbSet 2000 (printf "userCancel%d" .User.ID) $userCancelString}}
                    {{$report := index (getMessage $reportLog $reportMessageID).Embeds 0|structToSdict}}
                    {{range $k, $v := $report}}
                        {{if eq (kindOf $v true) "struct"}}
                            {{$report.Set $k (structToSdict $v)}}
                        {{end}}
                    {{end}}
                    {{$user := userArg (dbGet $reportMessageID "reportAuthor").Value}}
                    {{with $report}}
                        {{.Author.Set "Icon_URL" $report.Author.IconURL}} 
                        {{.Footer.Set "Icon_URL" $report.Footer.IconURL}}
                        {{.Set "description" $combinedString}}
                        {{.Set "color" 16711935}}
                        {{$.Set "Author" (sdict "text" (print $user.String "(ID" $user.ID ")") "icon_url" ($user.AvatarURL "256"))}}
                        {{.Set "Fields" ((cslice).AppendSlice .Fields)}}{{.Fields.Set 4 (sdict "name" "Reaction Menu Options" "value" $cancelGuide)}}
                    {{end}}
                    {{editMessage $reportLog $reportMessageID (complexMessageEdit "embed" $report)}}
                    Cancellation requested, have a nice day!
                    {{deleteAllMessageReactions $reportLog $reportMessageID}}
                    {{addMessageReactions $reportLog $reportMessageID "🚫" "✅" "⚠️"}}
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