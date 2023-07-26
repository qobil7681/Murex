import{_ as r}from"./plugin-vue_export-helper-c27b6911.js";import{r as s,o as d,c as i,d as t,b as a,w as n,e,f as l}from"./app-45f7c304.js";const u={},c=l(`<h1 id="tout" tabindex="-1"><a class="header-anchor" href="#tout" aria-hidden="true">#</a> <code>tout</code></h1><blockquote><p>Print a string to the STDOUT and set it&#39;s data-type</p></blockquote><h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2><p>Write parameters to STDOUT without a trailing new line character. Cast the output&#39;s data-type to the value of the first parameter.</p><h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2><pre><code>tout: data-type &quot;string to write&quot; -&gt; \`&lt;stdout&gt;\`
</code></pre><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2><pre><code>» tout: json { &quot;Code&quot;: 404, &quot;Message&quot;: &quot;Page not found&quot; } -&gt; pretty
{
    &quot;Code&quot;: 404,
    &quot;Message&quot;: &quot;Page not found&quot;
}
</code></pre><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2><p><code>tout</code> supports ANSI constants.</p><p>Unlike <code>out</code>, <code>tout</code> does not append a carriage return / line feed.</p><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>`,12),h=t("code",null,"(",-1),p=t("code",null,"cast",-1),m=t("code",null,"err",-1),_=t("code",null,"format",-1),f=t("code",null,"out",-1),g=t("code",null,"pretty",-1);function q(x,b){const o=s("RouterLink");return d(),i("div",null,[c,t("ul",null,[t("li",null,[a(o,{to:"/user-guide/ansi.html"},{default:n(()=>[e("ANSI Constants")]),_:1}),e(": Infixed constants that return ANSI escape sequences")]),t("li",null,[a(o,{to:"/commands/brace-quote.html"},{default:n(()=>[h,e(" (brace quote)")]),_:1}),e(": Write a string to the STDOUT without new line")]),t("li",null,[a(o,{to:"/commands/cast.html"},{default:n(()=>[p]),_:1}),e(": Alters the data type of the previous function without altering it's output")]),t("li",null,[a(o,{to:"/commands/err.html"},{default:n(()=>[m]),_:1}),e(": Print a line to the STDERR")]),t("li",null,[a(o,{to:"/commands/format.html"},{default:n(()=>[_]),_:1}),e(": Reformat one data-type into another data-type")]),t("li",null,[a(o,{to:"/commands/out.html"},{default:n(()=>[f]),_:1}),e(": Print a string to the STDOUT with a trailing new line character")]),t("li",null,[a(o,{to:"/commands/pretty.html"},{default:n(()=>[g]),_:1}),e(": Prettifies JSON to make it human readable")])])])}const T=r(u,[["render",q],["__file","tout.html.vue"]]);export{T as default};
