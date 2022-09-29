<main role="main" class="col-md-9 ml-sm-auto col-lg-10 px-md-4">
    <a class="btn btn-primary btn-add" href="/admin/add" role="button">添加</a>
    <div class="table-responsive">
        <table class="table table-striped table-sm">
            <thead>
            <tr>
                <th>ID</th>
                <th>标题</th>
                <th>简介</th>
                <th>内容</th>
                <th>图片</th>
                <th>分类</th>
                <th>标识</th>
                <th>排序</th>
                <th>浏览量</th>
                <th>评论数</th>
                <th>点赞数</th>
                <th>添加时间</th>
                <th>发布时间</th>
                <th>最后更新时间</th>
                <th>操作</th>
            </tr>
            </thead>
            <tbody>
            {{range $key,$a := .articles}}
            <tr>
                <td>{{$a.Id}}</td>
                <td>{{$a.Title}}</td>
                <td>{{$a.Introduction}}</td>
                <td>{{$a.Content}}</td>
                <td>{{$a.Image}}</td>
                <td>{{$a.CategoryId}}</td>
                <td>{{$a.Status}}</td>
                <td>{{$a.Sort}}</td>
                <td>{{$a.ViewCount}}</td>
                <td>{{$a.CommentCount}}</td>
                <td>{{$a.LikeCount}}</td>
                <td>{{$a.CreatedAt.Format "2006-01-02 15:04:05"}}</td>
                <td>{{$a.PublishedAt.Format "2006-01-02 15:04:05"}}</td>
                <td>{{$a.UpdatedAt.Format "2006-01-02 15:04:05"}}</td>
                <td>
                    <div class="btn-group" role="group" aria-label="Basic example">
                        <a role="button" class="btn btn-primary btn-min" href="/admin/{{$a.Id}}" >编辑</a>
                        <a role="button" class="btn btn-danger btn-min" onclick="delRow(this)" url="/admin/{{$a.Id}}" >删除</a>
                    </div>
                </td>
            </tr>
            {{end}}
            </tbody>
        </table>
    </div>
</main>