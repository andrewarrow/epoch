div p-0 
  {{ template "navbar" . }}
  div mt-9 space-y-9 text-center 
    div flex w-full justify-center 
      div text-2xl font-bold
        Projects Overview
  {{ $items := index . "items" }}
  {{ range $i, $item := $items }}
    {{ $name := index $item "name" }}
    div p-3 flex justify-center w-full 
      div w-1/2 bg-gray-700 py-3 rounded-lg hover:border-2 hover:cursor-pointer
        div flex 
          div
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="size-6"> <path stroke-linecap="round" stroke-linejoin="round" d="M12 6.75a.75.75 0 1 1 0-1.5.75.75 0 0 1 0 1.5ZM12 12.75a.75.75 0 1 1 0-1.5.75.75 0 0 1 0 1.5ZM12 18.75a.75.75 0 1 1 0-1.5.75.75 0 0 1 0 1.5Z" /> </svg>
          div
            {{ $name }}
  {{ end }}
