{{$introChan := 671173998301937694}}
{{$introMsg := 0}}
{{if (dbGet .User.ID "intro-msg")}}
{{$introMsg = toInt (dbGet .User.ID "intro-msg").Value}}
{{end}}
{{dbSet .User.ID "intro-msg" (toString .Message.ID)}} {{/* apparently integers get weird in db so store this as a string */}}
{{if (getMessage $introChan $introMsg)}}
     {{deleteMessage $introChan $introMsg}}
{{end}}
