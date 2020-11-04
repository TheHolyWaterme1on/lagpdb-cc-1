{{$introChan := 671173998301937694}}
{{$introMsg := 0}}
{{if (dbGet .User.ID "intro-msg")}}
{{$introMsg = toInt (dbGet .User.ID "intro-msg").Value}}
{{end}}
{{if (getMessage $introChan $introMsg)}}
	{{deleteMessage $introChan $introMsg}}
{{end}}
{{dbDel .User.ID "intro-msg"}}
