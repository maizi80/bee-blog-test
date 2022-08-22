
	<!-- 判断是否搜索 -->
	<div class="blank"></div>
	<div class="headertop">
		<!-- 首页大图 -->
		<div id="centerbg" style="background-image: url(/static/images/headerbg.jpg);">
			<!-- 左右倾斜 -->
			<div class="slant-left"></div>
			<div class="slant-right"></div>
			<!-- 博主信息 -->
			<div class="focusinfo">
				<!-- 头像 -->
				<div class="header-tou">
					<a href="/"><img src="/static/images/tx20110211.png"></a>
				</div>
				<!-- 简介 -->
				<div class="header-info">
					<p>Learning, recording, growth and life</p>
				</div>
				<!-- 社交信息 -->
				<ul class="top-social">
					<li class="qq">
						<a href="#"><img src="/static/images/qq.png"></a>
						<div class="qqInner">bee-blog</div>
					</li>
					<li class="qq">
						<a href="#" target="_blank" rel="nofollow noopener noreferrer"><img src="/static/images/coolapk.png"></a>
						<div class="qqInner">bee-blog</div>
					</li>
					<li><a href="#" target="_blank" rel="nofollow noopener noreferrer" class="social-github"><img src="/static/images/github.png"></a></li>
					<li><a href="#" target="_blank" rel="nofollow noopener noreferrer" class="social-bilibili"><img src="/static/images/bilibili.png"></a></li>
					<li><a href="#" target="_blank" rel="nofollow noopener noreferrer" class="social-bilibili"><img src="/static/images/music.png"></a></li>
				</ul>
			</div>
		</div>
		<!-- 首页大图结束 -->
	</div>
	<div class=""></div>
	<div id="content" class="site-content">
	<!-- 顶部公告内容 -->
	<div class="notice">
		<i class="iconfont"></i>
		<div class="notice-content">学习、记录、成长、生活</div>
	</div>
	<!-- 聚焦内容 -->
	<div class="top-feature">
		<h1 class="fes-title">聚焦</h1>
		<ul class="feature-content">
			<li class="feature-1"><a href="/Web/Akina.html"><div class="feature-title"><span class="foverlay">Akina</span></div><img src="/static/images/feature/feature1.jpg"></a></li>
			<li class="feature-2"><a href="/Web/userAkina.html"><div class="feature-title"><span class="foverlay">使用说明</span></div><img src="/static/images/feature/feature2.jpg"></a></li>
			<li class="feature-3"><a href="/archives.html"><div class="feature-title"><span class="foverlay">文章归档</span></div><img src="/static/images/feature/feature3.jpg"></a></li>
		</ul>
	</div>
	<!-- 主页内容 -->
	<div id="primary" class="content-area">
		<main id="main" class="site-main indexMain" role="main">
		<h1 class="main-title">近况</h1>
<!-- 结束搜索判断 -->
		<!-- 开始文章循环输出 -->
		{{range $key,$article := .articles}}
    	<article class="post post-list" itemscope="" itemtype="http://schema.org/BlogPosting">
		<!-- 判断文章输出样式 -->
			<div class="post-entry">
				<div class="feature">
					<a href="/Web/Akina.html"><div class="overlay"><i class="iconfont"></i></div>
      			<img src="{{$article.Image}}">
    			</a>
				</div>
				<h1 class="entry-title"><a href="/Web/Akina.html">{{if compare $article.Status 1}}<span style="color:#ff6d6d;font-weight:600">[置顶] </span>  {{end}}{{$article.Title}}</a></h1>
				<div class="p-time">
					<i class="iconfont"></i> {{date $article.PublishedAt "Y-m-d"}}<i class="iconfont hotpost" style="margin-left: 5px;"></i>
				</div>
				<a href="/Web/Akina.html"><p>{{$article.Introduction}}...</p></a>
				<!-- 文章下碎碎念 -->
				<footer class="entry-footer">
					<div class="post-more">
						<a href="/Web/Akina.html"><i class="iconfont"></i></a>
					</div>
					<div class="info-meta">
						<div class="comnum">
							<span><i class="iconfont"></i><a href="/Web/Akina.html">{{$article.CommentCount}}条评论</a></span>
						</div>
						<div class="views">
							<span><i class="iconfont"></i>{{$article.ViewCount}}万 热度</span>
						</div>
					</div>
				</footer>
			</div>
			<hr>
		</article>
		{{end}}

		<!-- 结束文章循环输出 -->
		<!-- 翻页按钮 -->
		<nav class="navigator">
			<i class="navnumber">1 / 9</i>
			<a class="next" title="" href="/page/2/"><i class="iconfont"></i></a>
		</nav>
	</main>
	<div id="pagination"><a class="next" title="" href="/page/2/">加载更多</a></div>
</div>
</div>
<!-- 结束主页内容 -->

