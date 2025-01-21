<svelte:head>
  <title>我公开的文章</title>
</svelte:head>
<script>
    import { page } from '$app/stores';
    import { goto } from '$app/navigation';
    import { browser } from '$app/environment';
    let articles = '';
    let data = '';
    let isLoading = true; 
    let error = '';
    let authoruid = $page.params.uid
    let authornickname = '';
    let search = '';
    let admin = false;

    let currentUserUid = '';
    function getUidFromCookies() {
        if(browser) {
             const cookies = document.cookie.split(';'); 
             for (let cookie of cookies) {
               cookie = cookie.trim(); 
              if (cookie.startsWith('uid=')) {
                return cookie.substring(4); 
            }
        }
         return null; 
        }
    }
    currentUserUid = getUidFromCookies();
    function ifcookie() {
      if(currentUserUid === '' || currentUserUid === null) {
        alert('请先登录！')
        goto('/')
      }
    }
    if(browser) ifcookie()
    async function Getadmin() {
      try {
        const res = await fetch(`/api/${currentUserUid}`);
        const data = await res.json();
        
        admin = data.user.admin
      } catch (error) {
        console.error('加载文章失败', error);
      }
    }
    if(browser) {
      Getadmin();
    } 

    async function fetchArticles() {
        try {
            const response = await fetch(`/api/articles/public/${authoruid}`); 
            if (!response.ok) {
                throw new Error('无法获取文章数据');
            }
            data = await response.json(); 
            articles = data.articles
            authornickname = data.authornickname
        } catch (err) {
           error = err.message;
        } finally {
            isLoading = false; 
        }
    }
    fetchArticles();

    function Maintext(text) {
        const richTextData = text

        // 创建一个临时 DOM 元素
        const tempDiv = document.createElement("div");
        tempDiv.innerHTML = richTextData;

        // 提取纯文本
        const plainText = tempDiv.innerText; // 或 tempDiv.textContent
    //    const previewText = plainText.slice(0, 50);
        
        return plainText;
    }

    
    function goToPublicArticles() {
      goto('/article/public');
     }
    function goTOMyArticles() {
        goto(`/article/${currentUserUid}/myarticle`)
    }
    function gotomypublicarticles(){
        goto(`/article/${currentUserUid}/public`)
    }
    function gotomyprivatearticles(){
        goto(`/article/${currentUserUid}/private`)
    }
    function gotocreatearticles(){
        goto(`/article/${currentUserUid}/create`)
    }
    function gotomypage(){
        goto(`/article/${currentUserUid}`)
    }
    function gotosearch(){
      if(search === '') {
        alert('搜索不能为空！')
        return
      }
      goto(`/search/${search}/article`)
    }
    function gotoviewart(articleid) {
      goto(`/article/view/${articleid}`)
    }
    function gotocreateuser() {
      goto(`/article/${currentUserUid}/createuser`)
    }
    function gotoprivatearticle() {
      goto(`/article/private`)
    }
    async function exit() {
      if (currentUserUid === null) {
        alert("请先登录")
        goto("/")
        return
      }
        try {
            const response = await fetch(`/api/exit/${currentUserUid}`,{
              method: 'POST',
            }); 
            if (!response.ok) {
                throw new Error('退出登录失败');
            }
            alert("退出登录成功")
            goto("/")
            
        } catch (err) {
           error = err.message;
        } finally {
            isLoading = false; 
        }
    }
</script>
  

<main>

  <div class="guide">
    <div class="search-container">
      <input type="text" id="search-input" class="search-input" placeholder="搜索..." bind:value={search}>
      <button type="submit" class="search-button" on:click={gotosearch}>
        <i class="search-icon">🔍</i>
      </button>
    </div>
  
    {#if admin}
    <button class="menu-button" on:click={gotocreateuser}>创建账号</button>
   {/if}
   {#if admin}
   <button class="menu-button" on:click={gotoprivatearticle}>未公开文章</button>
  {/if}
    <button class="menu-button" on:click={goToPublicArticles}>公开文章</button>
    <div class="menu">
      <button class="menu-button" on:click={goTOMyArticles}>我的文章</button>
      <div class="dropdown">
        <button class="dropdown-item"  on:click={gotomypublicarticles}>我公开的文章</button>
        <button class="dropdown-item"  on:click={gotomyprivatearticles}>未公开的文章</button>
        <button class="dropdown-item"  on:click={gotocreatearticles}>创建文章</button>
      </div>
    </div>
  
    <div class = "menu">
      <button class="menu-button" on:click={gotomypage}>我的主页</button>
      <div class = "dropdown">
        <button class="dropdown-item"  on:click={exit}>退出登录</button>
      </div>
  </div>
  </div>


    {#if isLoading} 
        <p>加载中...</p>
    {:else if error}
        <p style="color: red;">{error}</p>
    {:else}
      <div class = "article-con">
        <div class = "head">
            <h1>用户{authornickname}已公开的文章</h1>
        </div>
        {#if articles === null} 
          <p>没有文章</p>
        {/if}
            {#each articles as article}
                <button class="article" on:click={gotoviewart(article.id)}>
                   <h1>{article.title}</h1>
                    <p>作者: {article.author}</p>
                    <p>{Maintext(article.content)}</p>
                   <p>状态：已公开</p>
                   <p>创建时间：{article.createtime}</p>
                </button>
          
            {/each}
      </div>
    {/if}
    
</main>

<style>
    .guide {
        width: 100%;
        background-color: rgb(219, 220, 221);
        display: flex;
        justify-content: flex-end;
        gap:50px;
        margin-top: 0 ;
        padding-right: 5px;
    }
    .head{
        text-align: center;
    }

    .article {
    border: 2px solid rgb(114, 97, 97);
    border-radius: 10px;
    background: none; 
    color: inherit; 
    font: inherit;
    text-align: left; 
    padding: 0;
    margin: 0; 
    width: 100%; 
    cursor: pointer; 
}
.article h1, .article p {
    margin: 5px 0;           
    overflow: hidden;        
    text-overflow: ellipsis; 
    white-space: nowrap;    
}
.article:hover {
    background-color: rgba(0, 0, 0, 0.05); 
}

    .article-con{
        display: flex;
        flex-direction :column;
        gap: 10px;
    }
    .menu{
        position: relative;
    }
    .dropdown {
        position:absolute;
        top: 100%;
        left: -25%;
        display: none;
        background-color: rgb(219, 220, 221);
        white-space: nowrap;
    }
    .menu:hover .dropdown{
        display: block;
    }
    .search-container {
        display: flex;
        
    }
    .guide {
  width: 100%;
  background-color: rgb(219, 220, 221);
  display: flex;
  justify-content: flex-end;
  gap: 50px;
  margin-top: 0;
  padding-right: 5px;
}

.menu {
  position: relative;
}

.menu-button {
  background: none;
  border: none;
  font-size: inherit;
  cursor: pointer;
  padding: 10px 20px;
}

.dropdown {
  position: absolute;
  top: 100%;
  left: 0;
  display: none; /* 确保初始时隐藏 */
  flex-direction: column;
  align-items: stretch;
  background-color: rgb(219, 220, 221);
  white-space: nowrap;
  padding: 5px 0;
  box-shadow: 0px 4px 6px rgba(0, 0, 0, 0.1);
}

.dropdown-item {
  background: none;
  border: none;
  font-size: inherit;
  padding: 10px 15px;
  cursor: pointer;
  text-align: left;
  width: 100%;
}

.dropdown-item:hover {
  background-color: rgb(200, 200, 200);
}


.menu:hover .dropdown {
  display: flex; 
}

.dropdown:hover {
  display: flex; 
}
</style>