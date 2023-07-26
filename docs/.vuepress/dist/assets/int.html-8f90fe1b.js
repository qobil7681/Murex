import{_ as r}from"./plugin-vue_export-helper-c27b6911.js";import{r as l,o as i,c as s,d as e,b as n,w as a,e as t,f as d}from"./app-45f7c304.js";const u={},c=d('<h1 id="int-data-type-reference" tabindex="-1"><a class="header-anchor" href="#int-data-type-reference" aria-hidden="true">#</a> <code>int</code> - Data-Type Reference</h1><blockquote><p>Whole number (primitive)</p></blockquote><h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2><p>An integer is a whole number (eg 1, 2, 3, 4) rather than one with a decimal point (such as 1.1).</p><p>Integers in Murex are sized based on the bit (or word) size of the target CPU.</p><p>A 386, ARMv6 or other 32bit build of Murex would see the range of from <code>-2147483648</code> (negative) through <code>2147483647</code> (positive).</p><p>AMD64 or other 64bit built of Murex would see the range from <code>-9223372036854775808</code> (negative) through <code>9223372036854775807</code> (positive).</p><blockquote><p>Unless you specifically know you only want whole numbers, it is recommended that you use the default numeric data-type: <code>num</code>.</p></blockquote><h2 id="supported-hooks" tabindex="-1"><a class="header-anchor" href="#supported-hooks" aria-hidden="true">#</a> Supported Hooks</h2><ul><li><code>Marshal()</code> Supported</li><li><code>Unmarshal()</code> Supported</li></ul><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>',11),h=e("code",null,"Marshal()",-1),m=e("code",null,"Unmarshal()",-1),p=e("code",null,"[[",-1),f=e("code",null,"[",-1),_=e("code",null,"cast",-1),b=e("code",null,"format",-1),y=e("code",null,"num",-1),g=e("code",null,"open",-1),x=e("code",null,"runtime",-1),v=e("code",null,"str",-1);function k(w,M){const o=l("RouterLink");return i(),s("div",null,[c,e("ul",null,[e("li",null,[n(o,{to:"/apis/Marshal.html"},{default:a(()=>[h,t(" (type)")]),_:1}),t(": Converts structured memory into a structured file format (eg for stdio)")]),e("li",null,[n(o,{to:"/apis/Unmarshal.html"},{default:a(()=>[m,t(" (type)")]),_:1}),t(": Converts a structured file format into structured memory")]),e("li",null,[n(o,{to:"/commands/element.html"},{default:a(()=>[p,t(" (element)")]),_:1}),t(": Outputs an element from a nested structure")]),e("li",null,[n(o,{to:"/commands/index2.html"},{default:a(()=>[f,t(" (index)")]),_:1}),t(": Outputs an element from an array, map or table")]),e("li",null,[n(o,{to:"/commands/cast.html"},{default:a(()=>[_]),_:1}),t(": Alters the data type of the previous function without altering it's output")]),e("li",null,[n(o,{to:"/commands/format.html"},{default:a(()=>[b]),_:1}),t(": Reformat one data-type into another data-type")]),e("li",null,[n(o,{to:"/types/num.html"},{default:a(()=>[y,t(" (number)")]),_:1}),t(": Floating point number (primitive)")]),e("li",null,[n(o,{to:"/commands/open.html"},{default:a(()=>[g]),_:1}),t(": Open a file with a preferred handler")]),e("li",null,[n(o,{to:"/commands/runtime.html"},{default:a(()=>[x]),_:1}),t(": Returns runtime information on the internal state of Murex")]),e("li",null,[n(o,{to:"/types/str.html"},{default:a(()=>[v,t(" (string) ")]),_:1}),t(": string (primitive)")])])])}const C=r(u,[["render",k],["__file","int.html.vue"]]);export{C as default};
