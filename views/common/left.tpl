	<div class="col-md-4">
			<div class="page-header">
				<h1>
					 <small>welcome</small>
				</h1>
			</div>
 
			 <div class="list-group">
                    <a href="#" class="list-group-item disabled">分类</a> 
					{{range .types}}
                    <a href="/?tid={{.Id}}" class="list-group-item"><span class="badge">{{.Nums}}</span>{{.Name}}</a> 
					{{end}}
                </div>
				
			 
		</div>
