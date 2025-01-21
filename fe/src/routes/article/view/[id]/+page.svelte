<script lang="ts">
    import { onMount } from 'svelte';
    import { goto } from '$app/navigation'; 
    import { page } from '$app/stores';
    import { browser } from '$app/environment';
  
    let title = '';
    let articleid = ''; 
    let content = '';
    let authoruid = '';
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
    function gotocreateuser() {
      goto(`/article/${currentUserUid}/createuser`)
    }
    function gotoedit() {
      if (currentUserUid !== authoruid && !admin) {
        alert('你没有编辑权限！')
        return
      }
      goto(`/article/${currentUserUid}/create/${articleid}`)
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
  
    onMount(async () => {
      articleid = $page.params.id;
       
      if (articleid) {
        await loadArticle();
      }
    });
  
    // 加载文章
    async function loadArticle() {
      try {
        const res = await fetch(`/api/articles/${articleid}`);
        const data = await res.json();
        authoruid = data.authoruid
        title = data.title;
        content = data.content;
      } catch (error) {
        console.error('加载文章失败', error);
      }
    }
  </script>
  <svelte:head>
    <title>{title}</title>
  </svelte:head>
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

    <div class="container">
      <h1 class="head">{title}</h1>
      <!-- 显示文章内容 -->
      <div id="viewer" class = "viewer">{@html content}</div>
    </div>

    {#if currentUserUid === authoruid || admin}
      <div class="buttons">
       <button on:click={gotoedit}>编辑</button>
     </div>
    {/if}
    <p>文章ID：{articleid}</p>
    <p>作者UID：{authoruid}</p>
    
  </main>
  
  <style>
    .container {
      width: 80%;
      margin: 50px auto;
      border: 2px solid rgb(114, 97, 97);
      border-radius: 10px;
      padding: 20px;
      box-sizing: border-box;
    }
    .viewer {
  width: 100%;
  min-height: 300px;
  border: 1px solid #ccc;
  padding: 15px;
  background: #fff;
  box-sizing: border-box;
  word-wrap: break-word; /* 自动换行长单词 */
  overflow-wrap: break-word; /* 支持换行 */
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
    .search-container {
      display: flex;
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
    .head {
      text-align: center;
    }

    .guide {
        width: 100%;
        background-color: rgb(219, 220, 221);
        display: flex;
        justify-content: flex-end;
        gap:50px;
        margin-top: 0 ;
        padding-right: 5px;
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
    .head {
        text-align: center;
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
.buttons {
    display: flex;
    gap: 50px;
    justify-content: center;
}

.buttons button {
    padding: 20px 40px;    
    font-size: 18px;       
    border-radius: 10px;   
    background-color: #4CAF50; 
    color: white;           
    border: none;           
    cursor: pointer;      
    transition: all 0.3s ease; 
}

/* 鼠标悬浮时 */
.buttons button:hover {
    background-color: #45a049; 
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2); 
    transform: scale(1.05);
}


</style>
  