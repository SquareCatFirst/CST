<svelte:head>
  <title>编辑文章</title>
</svelte:head>
<script lang="ts">
  import { onMount } from 'svelte';
  import { writable } from 'svelte/store';
  import { goto } from '$app/navigation'; 
  import type Quill from 'quill'; 
  import { page } from '$app/stores';
  import { browser } from '$app/environment';

  let quill: Quill | undefined; 
  let title = '';
  let articleuid = ''; 
  let content = ''; 
  let isLoading = writable(false);
  let authoruid = '';
  let search = '';
  let admin = false;


  const wordLimit = 20000;
  const maxLength = 100;
  //判断文章是否超字数
function checkWordLimit() {
  if (quill) {
    
    const text = quill.getText();
    
    
    const textLength = text.length;

    
    if (textLength > wordLimit) {
      alert('超出字数限制！'+wordLimit+'以内,当前字数:'+textLength)
      return false;
    } else {
      return true;
    }
  }
}

//判空
function isEmptyOrWhitespace() {
  if (quill) {
    
    const text = quill.getText();
    
    
    const trimmedText = text.trim();
    
    if (trimmedText === "") {
      alert('无效内容！')
      return false;
    } else {
      return true;
    }
  }
}
function isTitleTooLong() {
  if (title.length > maxLength) {
    alert('标题过长，当前标题字数为:'+title.length+'限制长度为:'+maxLength)
    return false;
  } else {
    return true;
  }
}

function isTitleEmptyOrWhitespace() {
  //判空
  const trimmedTitle = title.trim();
  
  if (trimmedTitle === "") {
    alert('无效标题！')
    return false;
  } else {
    return true;
  }
}

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
    articleuid = $page.params.articleuid
    authoruid = $page.params.uid
    if(currentUserUid !== authoruid && !admin) {
      alert("你没有编辑权限！")
      goto('/article/public')
      return
    }
    //  quill
    if (typeof window !== 'undefined') {
      const { default: Quill } = await import('quill');
      quill = new Quill('#editor', {
      theme: 'snow',
      modules: {
        toolbar: [
          [{ 'header': '1'}, { 'header': '2'}, { 'font': [] }],
          [{ 'list': 'ordered'}, { 'list': 'bullet' }],
          ['bold', 'italic', 'underline', 'strike'],
          [{ 'align': [] }],
          ['link', 'image',],
          [{ 'color': [] }, { 'background': [] }], 
        ],
      }
    });

      if (articleuid) {
        await loadArticle();
      }
    }
  });

  // 加载文章
  async function loadArticle() {
    if (currentUserUid !== authoruid && !admin) {
        alert('你没有编辑权限！')
        return
      }
    try {
      const res = await fetch(`/api/articles/${articleuid}`);
      const data = await res.json();
      title = data.title;
      content = data.content;
      if (quill) {
        quill.root.innerHTML = content; 
      }
    } catch (error) {
      console.error('加载文章失败', error);
    }
  }

  // 保存文章
  async function saveArticle() {
    if (currentUserUid !== authoruid && !admin) {
        alert('你没有编辑权限！')
        return
      }
   
    isLoading.set(true);
    const articleData = {
      articleuid,
      authoruid,
      title,
      content: quill?.root.innerHTML || '', 
    };
    if(!checkWordLimit() || !isEmptyOrWhitespace() || !isTitleTooLong() || !isTitleEmptyOrWhitespace()) {
      return;
    }
    if(articleData.title === '' ){
          alert('不可保存无标题文章！')
          return
       }

    try {
      if (articleuid) {
        // 更新已有文章
        const res = await fetch(`/api/articles/${articleuid}`, {
          method: 'PUT',
          body: JSON.stringify(articleData),
          headers: { 'Content-Type': 'application/json' },
        });
        if (res.ok) {
          alert('保存成功');
          goto(`/article/${authoruid}/create/${articleuid}`);
        } else {
          const data = await res.json()
          alert('保存失败:'+data.message)
        }
      } else {
        // 创建新文章
        const res = await fetch('/api/articles', {
          method: 'POST',
          body: JSON.stringify(articleData),
          headers: { 'Content-Type': 'application/json' },
        });
        const data = await res.json();
        articleuid = data.uuid; // 获取新文章的 UUID
        if (res.ok) {
          alert('保存成功');
          goto(`/article/${authoruid}/create/${articleuid}`); 
        } else {
          const data = await res.json()
          alert('保存失败 :'+data.message)
        }
      }
      
    } catch (error) {
      console.error('保存文章失败', error);
      alert('不可保存空文章');
    } finally {
      isLoading.set(false);
    }
  }

  // 公开文章
  async function publishArticle() {
    if (currentUserUid !== authoruid && !admin) {
        alert('你没有编辑权限！')
        return
      }
      
    isLoading.set(true);
    const articleData = {
      articleuid,
      authoruid,
      title,
      content: quill?.root.innerHTML || '', // 确保 quill 被正确初始化
    };
    if(!checkWordLimit() || !isEmptyOrWhitespace() || !isTitleTooLong() || !isTitleEmptyOrWhitespace()) {
      return;
    }
    if(articleData.title === '' ){
          alert('不可保存无标题文章！')
          return
       }
    try {
      if (articleuid) {
        // 更新已有文章
        const res = await fetch(`/api/articles/${articleuid}`, {
          method: 'PUT',
          body: JSON.stringify(articleData),
          headers: { 'Content-Type': 'application/json' },
        });
        if (res.ok) {
          goto(`/article/${authoruid}/create/${articleuid}`);
        } else {
          alert('保存失败')
          return
        }
      } else {
        // 创建新文章
        const res = await fetch('/api/articles', {
          method: 'POST',
          body: JSON.stringify(articleData),
          headers: { 'Content-Type': 'application/json' },
        });
        const data = await res.json();
        articleuid = data.uuid; // 获取新文章的 UUID
        if (res.ok) {
          goto(`/article/${authoruid}/create/${articleuid}`); // 跳转到文章页面
        } else {
          alert('保存失败')
          return
        }
      }
      
    } catch (error) {
      alert('不可保存空文章');
      return
    } finally {
      isLoading.set(false);
    }
    try {
      await fetch(`/api/articles/${articleuid}/publish`, {
        method: 'POST',
      });
      alert('文章已公开');
    } catch (error) {
      console.error('公开文章失败', error);
      alert('公开失败');
    }
  }

  // 删除文章
  async function deleteArticle() {
    if (currentUserUid !== authoruid && !admin) {
        alert('你没有编辑权限！')
        return
      }
    if(articleuid === '') {
      if( confirm('文章未保存,是否离开？')){
        goto("/article/public")
      }
      return
    }  
    if (confirm('确定要删除这篇文章吗？')) {
      try {
        const res = await fetch(`/api/articles/${articleuid}`, {
          method: 'DELETE',
        });
        if (res.ok) {
          alert('文章已删除');
          goto('/article/public'); // 跳转到文章列表页面
        } else {
          alert('文章删除失败');
        }
      } catch (error) {
        console.error('删除文章失败', error);
        alert('删除失败');
      }
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

    <div class="container">
        <h1 class = "head">编辑文章</h1>
        <!-- 标题输入框 -->
        <input 
          type="text" 
          bind:value={title} 
          placeholder="请输入标题" 
          class = "title"
        />
        
        <!-- 富文本编辑器 -->
        <div id="editor"></div>
      
        <!-- 操作按钮 -->
        <div class="buttons">
          <button on:click={saveArticle} >保存</button>
          <button on:click={publishArticle}>公开</button>
          <button on:click={deleteArticle}>删除</button>
        </div>
      </div>
</main>
  

<style>
   @import 'quill/dist/quill.snow.css';
    #editor {
        width: 100%;
        height: 500px; 
        border: 1px solid #ccc; 
        box-sizing: border-box;  
    }
    .container {
        width: 80%;
        margin: 50px auto;
        border: 2px solid rgb(114, 97, 97);
        border-radius: 10px;
        padding: 20px;
        box-sizing: border-box;
     }
    .title {
        width: 100%; 
        padding: 10px; 
        margin-bottom: 20px;
        border: 1px solid #ccc;  
        box-sizing: border-box;  
    }
    .buttons {
        display: flex;
        gap: 50px;
        justify-content: center;
       
    }
    .buttons button {
     padding: 10px 20px;
     cursor: pointer;
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
  display: none; 
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
  
  
  