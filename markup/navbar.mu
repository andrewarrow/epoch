{{ define "navbar" }}
  div navbar bg-base-200
    div navbar-start 
      div btn btn-ghost
        a href=/
          h1 text-2xl font-bold
            epoch
    div navbar-center flex hidden md:block
      ul menu menu-horizontal
    div navbar-end text-2xl
      +
  {{ end }}
