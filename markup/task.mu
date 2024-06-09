div p-3
  {{ $name := index . "name" }}
  {{ $project := index . "project" }}
  div flex justify-center w-full 
    div w-1/2 bg-gray-700 py-3 rounded-lg hover:border-2 hover:cursor-pointer
      div flex mb-2
        div
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="size-6"> <path stroke-linecap="round" stroke-linejoin="round" d="M12 6.75a.75.75 0 1 1 0-1.5.75.75 0 0 1 0 1.5ZM12 12.75a.75.75 0 1 1 0-1.5.75.75 0 0 1 0 1.5ZM12 18.75a.75.75 0 1 1 0-1.5.75.75 0 0 1 0 1.5Z" /> </svg>
        div
          {{ $name }}
      div text-left ml-6 mt-1
        {{ if $project }}
          {{ $pname := index $project "name" }}
          {{ $color := index $project "color" }}
          span {{$color}} rounded-md p-1 text-sm
            {{ $pname }}
        {{ end }}
