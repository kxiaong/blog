{{ define "content" }}

<div style="margin-left:25%">
  <div class="w3-container w3-teal">
    <div class="container-fluid main-container" id="J_main-container">
      {{ range .arr "arr" }}
      <div class="row post-container-wrapper">
        <div class="col-md-6">
          <div class="post-container">
            <h2 class="post-title">
              <a href="/" rel="bookmark">{{ .title1 }}</a>
            </h2>
            <div class="meta-box">
              <span class="m-post-date">
                <i class="fa fa-calendar-o"></i> {{.date1}}
              </span>
              <div class="post-content post-expect">
                {{ .abstract1 }}
                <a class="more-link btn btn-primary btn-xs" href="/">阅读全文</a>
              </div>
              <div class="meta-box post-bottom-meta-box hidden-print">
                <span class="tag-links">
                  <i class="fa fa-tags" aria-hidden="true"></i>
                  <a href="/tags/" rel="tag">{{ .tags1 }}</a>
                </span>
              </div>
            </div>
          </div>
        </div>
        <div class="col-md-6">
          <div class="post-container">
            <h2 class="post-title">
              <a href="/" rel="bookmark">{{ .title2 }}</a>
            </h2>
            <div class="meta-box">
              <span class="m-post-date">
                <i class="fa fa-calendar-o"></i> {{ .date2 }}
              </span>
              <div class="post-content post-expect">
                {{ .abstract2 }}
                <a class="more-link btn btn-primary btn-xs" href="/">阅读全文</a>
              </div>
              <div class="meta-box post-bottom-meta-box hidden-print">
                <span class="tag-links">
                  <i class="fa fa-tags" aria-hidden="true"></i>
                  <a href="/tags/" rel="tag">{{ .tags2 }}</a>
                </span>
              </div>
            </div>
          </div>
        </div>

      </div>
      {{ end }}
    </div>
  </div>
</div>

{{ end }}
