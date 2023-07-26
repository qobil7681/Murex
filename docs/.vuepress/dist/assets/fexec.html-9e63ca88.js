import{_ as l}from"./plugin-vue_export-helper-c27b6911.js";import{r as n,o as c,c as d,d as e,b as a,w as i,e as t,f as r}from"./app-45f7c304.js";const s={},u=r('<h1 id="fexec" tabindex="-1"><a class="header-anchor" href="#fexec" aria-hidden="true">#</a> <code>fexec</code></h1><blockquote><p>Execute a command or function, bypassing the usual order of precedence.</p></blockquote><h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2><p><code>fexec</code> allows you to execute a command or function, bypassing the usual order of precedence.</p><h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2><pre><code>fexec: flag command [ parameters... ] -&gt; `&lt;stdout&gt;`\n</code></pre><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2><pre><code>fexec: private /source/builtin/autocomplete.alias\n</code></pre><h2 id="flags" tabindex="-1"><a class="header-anchor" href="#flags" aria-hidden="true">#</a> Flags</h2><ul><li><code>builtin</code> Execute a Murex builtin</li><li><code>function</code> Execute a Murex public function</li><li><code>help</code> Display help message</li><li><code>private</code> Execute a Murex private function</li></ul><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2><h3 id="order-of-precedence" tabindex="-1"><a class="header-anchor" href="#order-of-precedence" aria-hidden="true">#</a> Order of precedence</h3><p>There is an order of precedence for which commands are looked up:</p><ol><li><p><code>runmode</code>: this is executed before the rest of the script. It is invoked by the pre-compiler forking process and is required to sit at the top of any scripts.</p></li><li><p><code>test</code> and <code>pipe</code> functions also alter the behavior of the compiler and thus are executed ahead of any scripts.</p></li><li><p>private functions - defined via <code>private</code>. Private&#39;s cannot be global and are scoped only to the module or source that defined them. For example, You cannot call a private function directly from the interactive command line (however you can force an indirect call via <code>fexec</code>).</p></li><li><p>Aliases - defined via <code>alias</code>. All aliases are global.</p></li><li><p>Murex functions - defined via <code>function</code>. All functions are global.</p></li><li><p>Variables (dollar prefixed) which are declared via <code>global</code>, <code>set</code> or <code>let</code>. Also environmental variables too, declared via <code>export</code>.</p></li><li><p>globbing: however this only applies for commands executed in the interactive shell.</p></li><li><p>Murex builtins.</p></li><li><p>External executable files</p></li></ol><p>You can override this order of precedence via the <code>fexec</code> and <code>exec</code> builtins.</p><h3 id="compatibility-with-posix" tabindex="-1"><a class="header-anchor" href="#compatibility-with-posix" aria-hidden="true">#</a> Compatibility with POSIX</h3><p>For compatibility with traditional shells like Bash and Zsh, <code>builtin</code> is an alias to `fexec builtin</p><h2 id="synonyms" tabindex="-1"><a class="header-anchor" href="#synonyms" aria-hidden="true">#</a> Synonyms</h2><ul><li><code>fexec</code></li><li><code>builtin</code></li></ul><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>',20),h=e("code",null,"alias",-1),p=e("code",null,"autocomplete",-1),f=e("code",null,"bg",-1),m=e("code",null,"builtins",-1),x=e("code",null,"event",-1),b=e("code",null,"exec",-1),_=e("code",null,"fg",-1),v=e("code",null,"function",-1),g=e("code",null,"jobs",-1),y=e("code",null,"open",-1),k=e("code",null,"private",-1),w=e("code",null,"source",-1);function E(M,S){const o=n("RouterLink");return c(),d("div",null,[u,e("ul",null,[e("li",null,[a(o,{to:"/commands/alias.html"},{default:i(()=>[h]),_:1}),t(": Create an alias for a command")]),e("li",null,[a(o,{to:"/commands/autocomplete.html"},{default:i(()=>[p]),_:1}),t(": Set definitions for tab-completion in the command line")]),e("li",null,[a(o,{to:"/commands/bg.html"},{default:i(()=>[f]),_:1}),t(": Run processes in the background")]),e("li",null,[a(o,{to:"/commands/runtime.html"},{default:i(()=>[m]),_:1}),t(": Returns runtime information on the internal state of Murex")]),e("li",null,[a(o,{to:"/commands/event.html"},{default:i(()=>[x]),_:1}),t(": Event driven programming for shell scripts")]),e("li",null,[a(o,{to:"/commands/exec.html"},{default:i(()=>[b]),_:1}),t(": Runs an executable")]),e("li",null,[a(o,{to:"/commands/fg.html"},{default:i(()=>[_]),_:1}),t(": Sends a background process into the foreground")]),e("li",null,[a(o,{to:"/commands/function.html"},{default:i(()=>[v]),_:1}),t(": Define a function block")]),e("li",null,[a(o,{to:"/commands/fid-list.html"},{default:i(()=>[g]),_:1}),t(": Lists all running functions within the current Murex session")]),e("li",null,[a(o,{to:"/commands/open.html"},{default:i(()=>[y]),_:1}),t(": Open a file with a preferred handler")]),e("li",null,[a(o,{to:"/commands/private.html"},{default:i(()=>[k]),_:1}),t(": Define a private function block")]),e("li",null,[a(o,{to:"/commands/source.html"},{default:i(()=>[w]),_:1}),t(": Import Murex code from another file of code block")])])])}const R=l(s,[["render",E],["__file","fexec.html.vue"]]);export{R as default};
