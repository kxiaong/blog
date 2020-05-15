{{define "header"}}
  <header>
    <div class="row">
      <div class="twelve columns">
        <div class="logo"><a href="index.html"><img alt="" src="/static/images/logo.jpg"></a></div>
        <nav id="nav-wrap">
          <a class="mobile-btn" href="#nav-wrap" title="Show navigation">Show navigation</a>
          <a class="mobile-btn" href="#" title="Hide navigation">Hide navigation</a>

          <ul id="nav" class="nav">
            <li class="current"><a href="index.html">首页</a></li>
            <li><span><a href="blog.html">文章</a></span>
              <ul>
              <li><a href="blog.html">技术</a></li>
              <li><a href="single.html">杂谈</a></li>
              </ul>
            </li>
            <li><a href="about.html">关于我</a></li>
          </ul>
        </nav>
    </div>
  </header>
{{end}}
