div p-0 
  {{ template "navbar" . }}
  div mt-9 space-y-9 text-center 
    div flex w-full justify-center 
      div text-2xl font-bold
        Create Project
    div flex w-full justify-center 
      div w-1/2
        input w-full text-2xl autofocus=true type=text input input-primary placeholder=Project-Name
    div flex w-full justify-center 
      div w-1/2
        Theme Color
        input id=color1 data-jscolor={} value=#FFFFFF
