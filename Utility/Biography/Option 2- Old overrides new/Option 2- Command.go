{{$introChan := 671173998301937694}}
{{$introMsg := 0}}
{{if (dbGet .User.ID "intro-msg")}}
{{$introMsg = toInt (dbGet .User.ID "intro-msg").Value}}
{{end}}

{{if .ExecData}}
	{{if (getMessage $introChan $introMsg)}}
		{{deleteMessage $introChan $introMsg}}
	{{end}}
	{{dbDel .User.ID "intro-msg"}}
{{else}}
	{{if (getMessage $introChan $introMsg)}}
		{{sendDM (printf "You already have an introduction! You can send a new one if you delete that one, or you can edit your original.\nYou can find your introduction here: https://discord.com/channels/%d/%d/%d" .Guild.ID $introChan $introMsg)}}
		{{deleteTrigger 5}}
	{{else}}
		{{dbSet .User.ID "intro-msg" (toString .Message.ID)}} {{/* apparently integers get weird in db so store this as a string */}}
	{{end}}
{{end}}
