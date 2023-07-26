import{_ as d}from"./plugin-vue_export-helper-c27b6911.js";import{r as s,o as l,c as i,d as e,b as a,w as n,e as t,f as r}from"./app-45f7c304.js";const c={},h=r(`<h1 id="out" tabindex="-1"><a class="header-anchor" href="#out" aria-hidden="true">#</a> <code>out</code></h1><blockquote><p>Print a string to the STDOUT with a trailing new line character</p></blockquote><h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2><p>Write parameters to STDOUT with a trailing new line character.</p><h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2><pre><code>out: string to write -&gt; \`&lt;stdout&gt;\`
</code></pre><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2><pre><code>» out Hello, World!
Hello, World!
</code></pre><p>For compatibility with other shells, <code>echo</code> is also supported:</p><pre><code>» echo Hello, World!
Hello, World!
</code></pre><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2><p><code>out</code> / <code>echo</code> output as <code>string</code> data-type. This can be changed by casting (<code>cast</code>) or using the <code>tout</code> function.</p><h3 id="ansi-constants" tabindex="-1"><a class="header-anchor" href="#ansi-constants" aria-hidden="true">#</a> ANSI Constants</h3><p><code>out</code> supports ANSI constants.</p><h2 id="synonyms" tabindex="-1"><a class="header-anchor" href="#synonyms" aria-hidden="true">#</a> Synonyms</h2><ul><li><code>out</code></li><li><code>echo</code></li></ul><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>`,17),u=e("code",null,"(",-1),p=e("code",null,">>",-1),m=e("code",null,">",-1),_=e("code",null,"cast",-1),f=e("code",null,"err",-1),g=e("code",null,"pt",-1),b=e("code",null,"read",-1),x=e("code",null,"read",-1),y=e("code",null,"tout",-1),S=e("code",null,"tread",-1),w=e("code",null,"read",-1),T=e("em",null,"typed",-1);function N(D,W){const o=s("RouterLink");return l(),i("div",null,[h,e("ul",null,[e("li",null,[a(o,{to:"/user-guide/ansi.html"},{default:n(()=>[t("ANSI Constants")]),_:1}),t(": Infixed constants that return ANSI escape sequences")]),e("li",null,[a(o,{to:"/commands/brace-quote.html"},{default:n(()=>[u,t(" (brace quote)")]),_:1}),t(": Write a string to the STDOUT without new line")]),e("li",null,[a(o,{to:"/commands/greater-than-greater-than.html"},{default:n(()=>[p,t(" (append file)")]),_:1}),t(": Writes STDIN to disk - appending contents if file already exists")]),e("li",null,[a(o,{to:"/commands/greater-than.html"},{default:n(()=>[m,t(" (truncate file)")]),_:1}),t(": Writes STDIN to disk - overwriting contents if file already exists")]),e("li",null,[a(o,{to:"/commands/cast.html"},{default:n(()=>[_]),_:1}),t(": Alters the data type of the previous function without altering it's output")]),e("li",null,[a(o,{to:"/commands/err.html"},{default:n(()=>[f]),_:1}),t(": Print a line to the STDERR")]),e("li",null,[a(o,{to:"/commands/pt.html"},{default:n(()=>[g]),_:1}),t(": Pipe telemetry. Writes data-types and bytes written")]),e("li",null,[a(o,{to:"/commands/read.html"},{default:n(()=>[b]),_:1}),t(": "),x,t(" a line of input from the user and store as a variable")]),e("li",null,[a(o,{to:"/commands/tout.html"},{default:n(()=>[y]),_:1}),t(": Print a string to the STDOUT and set it's data-type")]),e("li",null,[a(o,{to:"/commands/tread.html"},{default:n(()=>[S]),_:1}),t(": "),w,t(" a line of input from the user and store as a user defined "),T,t(" variable (deprecated)")])])])}const I=d(c,[["render",N],["__file","out.html.vue"]]);export{I as default};
