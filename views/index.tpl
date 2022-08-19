<!DOCTYPE html>
<html lang="zh-CN">
<head>
	<meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
	<title itemprop="name">Bee-blog</title>
	<meta name="description" content="学习、记录、成长、生活">
	<meta name="keywords" content="学习、记录、成长、生活">
	<link rel="shortcut icon" href="/static/images/favicon.ico" type="image/x-icon">
	<link rel="stylesheet" href="/static/css/style.css" type="text/css">
	<link rel="stylesheet" href="/static/css/OwO.css" type="text/css">
	<link rel="stylesheet" href="/static/css/dark.css" type="text/css">  <!-- 个性化选项 CSS 代码 -->
	<link rel="stylesheet" href="/static/css/blog.css" type="text/css">  <!-- 个性化选项 CSS 代码 -->
  
</head>
<body class="home blog hfeed">
<!-- 加载动画 -->
	<section id="main-container">
	<div class="openNav">
		<div class="iconflat">	 
			<div class="icon"></div>
		</div>
		<!-- logo则显示 -->
		<div class="site-branding">
			<div class="site-title"><a href="#"><img src="/static/images/akina.png"></a></div>		  
		</div>
		<!-- logo 结束 -->
	</div>	
	 <!-- 主页面显示 -->
<div id="page" class="site wrapper">
	<header class="site-header" role="banner">
		<div class="site-top">
			<!-- logo则显示 -->
			<div class="site-branding">
			<div class="site-title"><a href="#"><img src="/static/images/akina.png"></a></div>		  
			</div>
			<!-- logo 结束 -->
			<div id="login-reg">
				<!-- 个人头像 -->
				<!-- 如果用户未登录 -->
				<div class="ex-login">
					<a href="/admin/" target="_top">
						<i class="iconfont"></i>
					</a>
				</div>
			</div>
			<!-- 搜索 -->
			<div class="searchbox">
				<i class="iconfont js-toggle-search iconsearch"></i>
			</div> 
			<!-- 分类 -->
			<div class="lower">
				<nav>
					<ul class="menu">
						<li class="current-menu-item"><a href="#">首页</a></li>
						<li><a href="#">分类</a>
							<ul class="sub-menu">
								<li><a href="/category/Android/">Android</a></li>
								<li><a href="/category/Z-Turn/">Z-Turn</a></li>
								<li><a href="/category/Web/">Web</a></li>
								<li><a href="/category/Life/">Life</a></li>
							</ul>
						</li>
						<li><a href="/#">更多</a>
							<ul class="sub-menu">
								<li><a href="/message.html">留言</a></li>
								<li><a href="/tags.html">标签云</a></li>
								<li><a href="/archives.html">文章归档</a></li>
								<li><a href="/about.html">关于博主</a></li>
								<li><a href="/links.html">左邻右舍</a></li>
							</ul>
						</li>
					</ul>
					<!-- 隐藏后菜单图标 -->
					<i class="iconfont show-nav"></i>
				</nav>
			</div>
		</div>
		<!-- 到顶部按钮 -->
		<div class="cd-top-box">
			<a href="/#" class="cd-top"></a>
		</div>
	</header>
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
    	<article class="post post-list" itemscope="" itemtype="http://schema.org/BlogPosting">
		<!-- 判断文章输出样式 -->
			<div class="post-entry">
				<div class="feature">
					<a href="/Web/Akina.html"><div class="overlay"><i class="iconfont"></i></div>
      			<img src="/static/images/random/Akina.jpg">
    			</a>
				</div>
				<h1 class="entry-title"><a href="/Web/Akina.html"><span style="color:#ff6d6d;font-weight:600">[置顶] </span>Akina for Typecho 主题模板</a></h1>
				<div class="p-time">
					<i class="iconfont"></i> 2018-10-15<i class="iconfont hotpost" style="margin-left: 5px;"></i>
				</div>
				<a href="/Web/Akina.html"><p>简介Akina for Typecho除了原本Akina的特性外，加入了DNS预解析，或许可以加快站点的访问速度。支持了无插件的代码高...</p></a>
				<!-- 文章下碎碎念 -->
				<footer class="entry-footer">
					<div class="post-more">
						<a href="/Web/Akina.html"><i class="iconfont"></i></a>
					</div>
					<div class="info-meta">
						<div class="comnum">
							<span><i class="iconfont"></i><a href="/Web/Akina.html">473条评论</a></span>
						</div>
						<div class="views">
							<span><i class="iconfont"></i>13.2万 热度</span>
						</div>
					</div>
				</footer>
			</div>
			<hr>
		</article>
				<article class="post post-list" itemscope="" itemtype="http://schema.org/BlogPosting">
		<!-- 判断文章输出样式 -->
					<div class="post-entry">
				<div class="feature">
					<a href="/Web/WeCenterUpdate.html"><div class="overlay"><i class="iconfont"></i></div>
            <img src="/static/images/random/deu1.jpg">
          </a>
				</div>
				<h1 class="entry-title"><a href="/Web/WeCenterUpdate.html">WeCenter 4.0 全新安装与升级过程</a></h1>
				<div class="p-time">
				<i class="iconfont"></i> 2022-6-15				</div>
				<a href="/Web/WeCenterUpdate.html"><p>前言前写日子帮一个客户升级WeCenter，还是beta版本，bug有些多，好在官方迭代速度很快，过了几天遇到的大部分问题都解决了，遂...</p></a>
				<!-- 文章下碎碎念 -->
				<footer class="entry-footer">
					<div class="post-more">
							<a href="/Web/WeCenterUpdate.html"><i class="iconfont"></i></a>
					</div>
							<div class="info-meta">
						<div class="comnum">
							<span><i class="iconfont"></i><a href="/Web/WeCenterUpdate.html">1条评论</a></span>
						</div>
						<div class="views">
							<span><i class="iconfont"></i>428 热度</span>
						</div>
					</div>
				</footer>
			</div>
				<hr>
		</article>
				<article class="post post-list" itemscope="" itemtype="http://schema.org/BlogPosting">
		<!-- 判断文章输出样式 -->
					<div class="post-entry">
				<div class="feature">
					<a href="/Web/WeCenter-customad.html"><div class="overlay"><i class="iconfont"></i></div>
            <img src="/static/images/random/deu2.jpg">
          </a>
				</div>
				<h1 class="entry-title"><a href="/Web/WeCenter-customad.html">WeCenter 自定义广告插件</a></h1>
				<div class="p-time">
				<i class="iconfont"></i> 2022-6-11				</div>
				<a href="/Web/WeCenter-customad.html"><p>介绍本插件可以在侧边栏、内容开头结尾、页面开头等几个地方添加自定义的广告，及自行决定广告的图片与链接的广告地址。长条形广告长宽比尽量接...</p></a>
				<!-- 文章下碎碎念 -->
				<footer class="entry-footer">
					<div class="post-more">
							<a href="/Web/WeCenter-customad.html"><i class="iconfont"></i></a>
					</div>
							<div class="info-meta">
						<div class="comnum">
							<span><i class="iconfont"></i><a href="/Web/WeCenter-customad.html">NOTHING</a></span>
						</div>
						<div class="views">
							<span><i class="iconfont"></i>356 热度</span>
						</div>
					</div>
				</footer>
			</div>
				<hr>
		</article>
				<article class="post post-list" itemscope="" itemtype="http://schema.org/BlogPosting">
		<!-- 判断文章输出样式 -->
					<div class="post-entry">
				<div class="feature">
					<a href="/Web/Typecho-gravatar.html"><div class="overlay"><i class="iconfont"></i></div>
            <img src="/static/images/random/deu3.jpg">
          </a>
				</div>
				<h1 class="entry-title"><a href="/Web/Typecho-gravatar.html">解决 Typecho 后台头像被墙的问题</a></h1>
				<div class="p-time">
				<i class="iconfont"></i> 2022-6-9				</div>
				<a href="/Web/Typecho-gravatar.html"><p>起因Typecho 默认使用的是 gravatar 的头像，这个东西会根据用户的邮箱去调用你的头像，有两个问题：一是这个需要用户去主动...</p></a>
				<!-- 文章下碎碎念 -->
				<footer class="entry-footer">
					<div class="post-more">
							<a href="/Web/Typecho-gravatar.html"><i class="iconfont"></i></a>
					</div>
							<div class="info-meta">
						<div class="comnum">
							<span><i class="iconfont"></i><a href="/Web/Typecho-gravatar.html">2条评论</a></span>
						</div>
						<div class="views">
							<span><i class="iconfont"></i>340 热度</span>
						</div>
					</div>
				</footer>
			</div>
				<hr>
		</article>
				<article class="post post-list" itemscope="" itemtype="http://schema.org/BlogPosting">
		<!-- 判断文章输出样式 -->
					<div class="post-entry">
				<div class="feature">
					<a href="/Web/CommentNotifier.html"><div class="overlay"><i class="iconfont"></i></div>
            <img src="/static/images/random/deu4.jpg">
          </a>
				</div>
				<h1 class="entry-title"><a href="/Web/CommentNotifier.html">又一个新的邮件通知插件-CommentNotifier</a></h1>
				<div class="p-time">
				<i class="iconfont"></i> 2022-6-9				</div>
				<a href="/Web/CommentNotifier.html"><p>又双叒叕失效了，这已经是我Typecho使用过程中的第三个插件了，前两个但凡能发出去一封邮件我就不换了。翻来翻去还得看 泽泽，找到了 ...</p></a>
				<!-- 文章下碎碎念 -->
				<footer class="entry-footer">
					<div class="post-more">
							<a href="/Web/CommentNotifier.html"><i class="iconfont"></i></a>
					</div>
							<div class="info-meta">
						<div class="comnum">
							<span><i class="iconfont"></i><a href="/Web/CommentNotifier.html">NOTHING</a></span>
						</div>
						<div class="views">
							<span><i class="iconfont"></i>254 热度</span>
						</div>
					</div>
				</footer>
			</div>
				<hr>
		</article>
				<article class="post post-list" itemscope="" itemtype="http://schema.org/BlogPosting">
		<!-- 判断文章输出样式 -->
					<div class="post-entry">
				<div class="feature">
					<a href="/Android/337.html"><div class="overlay"><i class="iconfont"></i></div>
            <img src="/static/images/random/deu5.jpg">
          </a>
				</div>
				<h1 class="entry-title"><a href="/Android/337.html">基于 Debain 部署家庭服务 adguardhome、home-assistant。</a></h1>
				<div class="p-time">
				<i class="iconfont"></i> 2022-4-30				</div>
				<a href="/Android/337.html"><p>起初是用的 Rocky 搭建 的，刚开始用起来还好，慢慢的时间长了。我发现 Adguard 的处理速度一直保持在10ms左右就是下不去...</p></a>
				<!-- 文章下碎碎念 -->
				<footer class="entry-footer">
					<div class="post-more">
							<a href="/Android/337.html"><i class="iconfont"></i></a>
					</div>
							<div class="info-meta">
						<div class="comnum">
							<span><i class="iconfont"></i><a href="/Android/337.html">NOTHING</a></span>
						</div>
						<div class="views">
							<span><i class="iconfont"></i>537 热度</span>
						</div>
					</div>
				</footer>
			</div>
				<hr>
		</article>
				<article class="post post-list" itemscope="" itemtype="http://schema.org/BlogPosting">
		<!-- 判断文章输出样式 -->
					<div class="post-entry">
				<div class="feature">
					<a href="/Z-Turn/336.html"><div class="overlay"><i class="iconfont"></i></div>
            <img src="/static/images/random/deu6.jpg">
          </a>
				</div>
				<h1 class="entry-title"><a href="/Z-Turn/336.html">rsync+inotify 同步两个文件夹内容</a></h1>
				<div class="p-time">
				<i class="iconfont"></i> 2022-3-1				</div>
				<a href="/Z-Turn/336.html"><p>在 Akina 模板的开发过程中开发目录与项目运行目录不一样，需要一个工具时时同步两个文件，尝试过使用软连接，然而 php 不识别。百...</p></a>
				<!-- 文章下碎碎念 -->
				<footer class="entry-footer">
					<div class="post-more">
							<a href="/Z-Turn/336.html"><i class="iconfont"></i></a>
					</div>
							<div class="info-meta">
						<div class="comnum">
							<span><i class="iconfont"></i><a href="/Z-Turn/336.html">2条评论</a></span>
						</div>
						<div class="views">
							<span><i class="iconfont"></i>679 热度</span>
						</div>
					</div>
				</footer>
			</div>
				<hr>
		</article>
				<article class="post post-list" itemscope="" itemtype="http://schema.org/BlogPosting">
		<!-- 判断文章输出样式 -->
					<div class="post-entry">
				<div class="feature">
					<a href="/Z-Turn/335.html"><div class="overlay"><i class="iconfont"></i></div>
            <img src="/static/images/random/deu7.jpg">
          </a>
				</div>
				<h1 class="entry-title"><a href="/Z-Turn/335.html">Online 服务器-转移服务到家里服务器</a></h1>
				<div class="p-time">
				<i class="iconfont"></i> 2022-2-28				</div>
				<a href="/Z-Turn/335.html"><p>博客自从上次维护之后（维护记录博客维护记录）服务器就搬到了阿里云的轻量云上。一直对轻量云不太满意，恰好为了远程开发组了一台 Pve，就...</p></a>
				<!-- 文章下碎碎念 -->
				<footer class="entry-footer">
					<div class="post-more">
							<a href="/Z-Turn/335.html"><i class="iconfont"></i></a>
					</div>
							<div class="info-meta">
						<div class="comnum">
							<span><i class="iconfont"></i><a href="/Z-Turn/335.html">NOTHING</a></span>
						</div>
						<div class="views">
							<span><i class="iconfont"></i>780 热度</span>
						</div>
					</div>
				</footer>
			</div>
				<hr>
		</article>
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
</div>
</section>
<!-- 页底信息 -->
<footer id="colophon" class="site-footer" role="contentinfo">
	<!-- 请尊重作者！至少保留主题名称及其超链接，谢谢！ -->
	<div id="footer" class="site-info">
		Copyright © 2022 by <a href="https://zhebk.cn/" target="_blank" rel="nofollow">纸盒博客</a> - All rights reserved<span class="sep"> | </span>Theme : <a href="/Web/Akina.html" target="_blank" rel="nofollow">Akina For Typecho</a>
		<div class="footertext">
			<p><a href="https://beian.miit.gov.cn/" target="_blank" rel="nofollow noopener noreferrer">ICP备 xxxx号-1</a></p>

			<a target="_blank" href="#" style="display:inline-block;" rel="nofollow noopener noreferrer">
			    <img src="/static/images/gongan.png" style="float: left;height: 17px;">
			    <p style="float:left;height:20px;line-height:20px;margin: 0px 0px 0px 5px;">公网安备 xxx 号</p>
			</a>
		</div>
	</div>
</footer>
<div id="mo-nav">
	<div class="m-avatar">
		<img src="/static/images/tx20110211.png">
	</div>
	<div class="m-search">
		<form class="m-search-form" method="post" action="/" role="search">
			<input class="m-search-input" type="search" name="s" placeholder="搜索...">
		</form>
	</div>
	<ul id="nav_menu" class="menu">
		<li class="current-menu-item"><a href="/">首页</a></li>
		<li>
			<a href="/#">分类</a>
			<ul class="sub-menu">
				<li><a href="/category/Android/">Android</a></li>
				<li><a href="/category/Z-Turn/">Z-Turn</a></li>
				<li><a href="/category/Web/">Web</a></li>
				<li><a href="/category/Life/">Life</a></li>			
			</ul>
		</li>
		<li>
			<a href="/#">更多</a>
			<ul class="sub-menu">
				<li><a href="/message.html">留言</a></li>
				<li><a href="/tags.html">标签云</a></li>
				<li><a href="/DNS.html">纸盒DNS</a></li>
				<li><a href="/archives.html">文章归档</a></li>
				<li><a href="/about.html">关于博主</a></li>
				<li><a href="/links.html">左邻右舍</a></li>
			</ul>
		</li>
	</ul>
</div>
<!-- 搜索 -->
<form class="js-search search-form search-form--modal" method="post" action="/" role="search">
	<div class="search-form__inner">
		<div class="search-div">
			<h1 class="micromb-search">你想搜索什么...</h1>
			<i class="iconfont"></i>
			<input class="submit" type="search" name="s" placeholder="Search...">
		</div>
	</div>
</form>
<!-- 搜索结束 -->
<script type="text/javascript">
	var app = {"pjax":""};
	if (!!window.ActiveXObject || "ActiveXObject" in window) { //is IE?
	  alert('请抛弃万恶的IE系列浏览器吧。');
	}
	var xl = 0;
  	var transparent = 1;
</script>

<script type="text/javascript" src="/static/js/jquery.min.js."></script>
<script type="text/javascript" src="/static/js/jquery.preloader.js."></script>
<script type="text/javascript" src="/static/js/jquery.pjax.js."></script>
<script type="text/javascript" src="/static/js/baguetteBox.min.js."></script>
<script type="text/javascript" src="/static/js/global.js."></script>

<div role="dialog" id="baguetteBox-overlay">
	<div id="baguetteBox-slider"></div>
	<button type="button" id="previous-button" aria-label="Previous" class="baguetteBox-button">
		<svg width="44" height="60">
			<polyline points="30 10 10 30 30 50" stroke="rgba(255,255,255,0.5)" stroke-width="4" stroke-linecap="butt" fill="none" stroke-linejoin="round"></polyline>
		</svg>
	</button>
	<button type="button" id="next-button" aria-label="Next" class="baguetteBox-button">
		<svg width="44" height="60">
			<polyline points="14 10 34 30 14 50" stroke="rgba(255,255,255,0.5)" stroke-width="4" stroke-linecap="butt" fill="none" stroke-linejoin="round"></polyline>
		</svg>
	</button>
	<button type="button" id="close-button" aria-label="Close" class="baguetteBox-button">
		<svg width="30" height="30">
			<g stroke="rgb(160,160,160)" stroke-width="4">
				<line x1="5" y1="5" x2="25" y2="25"></line>
				<line x1="5" y1="25" x2="25" y2="5"></line>
			</g>
		</svg>
	</button>
</div>
<script type="text/javascript" src="/static/js/prism.js."></script>
<script type="text/javascript" src="/static/js/SmoothScroll.js."></script>

</body>
</html>