<!doctype html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <meta name="description" content="">
    <title>后台首页-Bee-Blog</title>
    <!-- Bootstrap core CSS -->
    <link href="/static/css/bootstrap.min.css" rel="stylesheet">
    <style>
        .bd-placeholder-img {
            font-size: 1.125rem;
            text-anchor: middle;
            -webkit-user-select: none;
            -moz-user-select: none;
            -ms-user-select: none;
            user-select: none;
        }

        @media (min-width: 768px) {
            .bd-placeholder-img-lg {
                font-size: 3.5rem;
            }
        }
    </style>
    <!-- Custom styles for this template -->
    <link href="/static/css/dashboard.css" rel="stylesheet">
</head>
<body>

<nav class="navbar navbar-dark sticky-top bg-dark flex-md-nowrap p-0 shadow">
    <a class="navbar-brand col-md-3 col-lg-2 mr-0 px-3" href="/admin">Bee-Blog</a>
    <button class="navbar-toggler position-absolute d-md-none collapsed" type="button" data-toggle="collapse" data-target="#sidebarMenu" aria-controls="sidebarMenu" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
    </button>
    <ul class="navbar-nav px-3">
        <li class="nav-item text-nowrap">
            <a class="nav-link" href="/signout">Sign out</a>
        </li>
    </ul>
</nav>

<div class="container-fluid">
    <div class="row">
        <nav id="sidebarMenu" class="col-md-3 col-lg-2 d-md-block bg-light sidebar collapse">
            <div class="sidebar-sticky pt-3">
                <ul class="nav flex-column">
                    <li class="nav-item"><a class="nav-link active" href="/admin"><span data-feather="home"></span>文章管理 </a></li>
                    <li class="nav-item"><a class="nav-link" href="/category"><span data-feather="home"></span>文章分类管理 </a></li>
                    <li class="nav-item"><a class="nav-link" href="/tags"><span data-feather="home"></span>文章标签管理 </a></li>
                    <li class="nav-item"><a class="nav-link" href="/profile"><span data-feather="file"></span>个人信息</a></li>
                </ul>
            </div>
        </nav>

        {{.LayoutContent}}
    </div>
</div>
<script src="/static/js/jquery.min.js" crossorigin="anonymous"></script>
<script src="/static/js/bootstrap.bundle.min.js"></script>

</body>
</html>
