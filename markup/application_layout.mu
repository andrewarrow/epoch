html data-theme=sunset
  head
    {{ $build := index . "build" }}
    {{ $og := index . "og" }}
    meta property=og:image content={{$og}}
    link rel=apple-touch-icon href=/assets/images/book.png
    link rel=apple-touch-startup-image href=/assets/images/book.png
    link rel=icon href=/assets/images/book.png
    link rel=stylesheet type=text/css href=/assets/css/tail.min.css?id!{{$build}}
    link rel=stylesheet type=text/css href=/assets/css/main.css?id!{{$build}}
    {{ if index . "USE_LIVE_TEMPLATES" }}
      script src=https://cdn.tailwindcss.com
    {{ end }}
    script src=/assets/javascript/wasm_exec.js?id!{{$build}}
    <script src="https://cdnjs.cloudflare.com/ajax/libs/jscolor/2.5.2/jscolor.min.js" integrity="sha512-qFhMEJrjI50TwLDGZ7Oi0ksTSWnFOqTNXhlqqUgWnE65S23rWUtQOv+tMNEybkMYSXKgAc3eg/SzkX+qrtJT/g==" crossorigin="anonymous" referrerpolicy="no-referrer"></script>
    script
      function $(id) { return document.getElementById(id); }
    title
      {{ index . "title" }}
    {{ index . "viewport" }}
  body
    div hidden2 absolute bottom-0 left-1/2 transform -translate-x-1/2 w-1/3 h-24 bg-gray-100 rounded-lg flex justify-center items-center text-black
      div ml-6 mr-6 text-green-600 text-4xl
        +
      div mr-9
        span
          Created project 
        span font-bold
          werfwfw.
        span
          You can select it from the menu on the top left.
    div id=flash bg-red-600 text-white text-center fixed top-0 left-0 w-full
      {{ index . "flash" }}
    div overflow-x-auto pl-3 pr-3 min-h-screen font-montserrat text-base
      {{ index . "content" }}
    div 
      div pb-32 footer items-center p-10 bg-base-200 text-base-content rounded
        div items-center grid-flow-col
          Copyright &copy; 2024 - All right reserved by epoch 
        div grid-flow-col gap-4 md:place-self-center md:justify-self-end
          a href=/space/about-us link link-hover
            About Us
          a href=/space/pricing link link-hover
            Pricing
          a href=/space/terms link link-hover
            Terms & Conditions
          a href=/space/privacy link link-hover
            Privacy Policy
    {{ index . "wasm" }}
