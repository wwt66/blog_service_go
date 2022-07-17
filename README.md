# 程序主要接口
本后端程序主要是给前端提供文章存储和标签存储的接口，同时还会管理文章标签的对应关系。

一篇文章可能有多个标签。

# 表设计
1. 文章表：包含主键、创建人、创建日期、修改人、修改日期、删除日期、是否删除。标题名、描述、图片地址、内容、是否禁用
2. 标签表：包含主键、创建人、创建日期、修改人、修改日期、删除日期、是否删除。标签名、是否禁用
3. 文章标签关联表：文章和标签通过主键关联，包含主键、创建人、创建日期、修改人、修改日期、删除日期、是否删除。以及文章和标签的主键

# 接口设计
1. 文章相关：
   1. 获取多个文章：根据文章名、标签ID、状态与否过滤并得到所有文章，并通过分页获得第n页的文章
   2. 获取单个文章：根据文章主键获取文章
   3. 创建文章：根据标签ID、文章的其他信息创建一篇文章
   4. 更新文章：根据标签ID等信息更新一篇文章
   5. 删除文章：根据文章ID删除一篇文章
2. 标签相关：
   1. 获取多个标签：根据标签名、状态信息过滤得到标签，可以使用分页
   2. 创建标签：根据标签名，状态和创建者创建一个标签
   3. 更新标签：根据主键、标签信息和修改者修改一个标签
   4. 删除标签：根据标签主键删除一个标签