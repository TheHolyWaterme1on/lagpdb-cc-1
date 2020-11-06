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
    {{$reportMessage := (getMessage $reportLog $reportMessageID)}}
    {{if eq (toInt64 (dbGet $reportMessageID "reportAuthor").Value) (toInt64 .User.ID)}}
            {{if eq "used" $userKey}}
                Your latest report already has been cancelled!
            {{else}}
            {{if eq (index .CmdArgs 1|str) $userKey}}
                {{if ge (len .CmdArgs) 3}}
                    {{$reason := joinStr " " (slice .CmdArgs 2)}}
                    {{$userReportString := (dbGet 2000 (printf "userReport%d" .User.ID)).Value|str}}
                    {{$cancelGuide := (printf "Deny request with 🚫, accept with ✅, or request more information with ⚠️")}}
                    {{dbSet 2000 "cancelGuideBasic" $cancelGuide}}
                    {{$userCancelString := (printf "<@%d> requested cancellation of this report due to: `%s`" .User.ID $reason)}}
                    {{$combinedString := (print $userReportString " \n " $userCancelString)}}
                    {{dbSet 2000 (printf "userCancel%d" .User.ID) $userCancelString}}
                    {{if $reportMessage.Embeds}} {{/*structToSdict the embed so that we can edit it*/}}
                        {{$embed := structToSdict (index $reportMessage.Embeds 0)}}
                            {{range $k, $v := $embed}}
                                {{- if eq (kindOf $v true) "struct" }}
                                    {{- $embed.Set $k (structToSdict $v) }}
                                {{- end -}}
                            {{end}}
                        {{if $embed.Author}} {{$embed.Author.Set "Icon_URL" $embed.Author.IconURL}} {{end}}
                        {{if $embed.Footer}} {{$embed.Footer.Set "Icon_URL" $embed.Footer.IconURL}} {{end}}
                        {{$embed.Set "Description" $combinedString}}
                    {{end}}
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