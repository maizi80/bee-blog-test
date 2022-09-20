<main role="main" class="col-md-9 ml-sm-auto col-lg-10 px-md-4">

    <h2><a href="/category">添加</a></h2>
    <div class="table-responsive">
        <table class="table table-striped table-sm">
            <thead>
            <tr>
                <th>ID</th>
                <th>别名</th>
                <th>名字</th>
                <th>排序</th>
                <th>操作</th>
            </tr>
            </thead>
            <tbody>
            {{range $key,$category := .categorys}}
            <tr>
                <td>{{$category.Id}}</td>
                <td>{{$category.Alias}}</td>
                <td>{{$category.Name}}</td>
                <td>{{$category.Sort}}</td>
                <td>编辑|删除</td>
            </tr>
            {{end}}
            </tbody>
        </table>
    </div>
</main>