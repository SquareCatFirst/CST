<script>
	import { page } from '$app/stores';
    import { goto } from '$app/navigation';
    import { browser } from '$app/environment';
    let user = '';
    let data = '';
    let isLoading = true; 
    let error = '';
    let uid  = $page.params.uid;

    let currentUserUid = '';
	let isEditingPassword = false;
	let newPassword = '';
	let confirmPassword = '';
  let search = '';
    let admin = '';

    // 获取cookie中的uid
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

    async function fetchArticles() {
        try {
            const response = await fetch(`/api/${uid}`); 
            if (!response.ok) {
                throw new Error('无法获取用户数据');
            }
            data = await response.json(); 
            user = data.user
            
        } catch (err) {
           error = err.message;
        } finally {
            isLoading = false; 
        }
    }

    // 修改密码提交
	async function changePassword() {
		if (newPassword !== confirmPassword) {
			alert('两次密码输入不一致');
			return;
		}
    if (newPassword === '' && confirmPassword === '') {
      alert('密码不能为空！')
      return
    }

		try {
			const response = await fetch(`/api/users/password/${uid}`, {
				method: 'PUT',
				headers: {
					'Content-Type': 'application/json',
				},
				body: JSON.stringify({ password: newPassword }),
			});
			if (!response.ok) {
				throw new Error('修改密码失败');
			}
			alert('密码修改成功');
			isEditingPassword = false;
            window.location.reload();
		} catch (err) {
			alert(err.message);
		}
	}

    fetchArticles();
   /*
    如果当前用户是管理员，则可以在该用户的主页去看该用户所有文章，公开文章和未公开的文章，与修改密码
    如果当前用户是主页用户，则可以在自己的主页去选择看自己所有文章，公开的文章与未公开的文章，与修改密码
    如果当前用户不是主页用户，则只可以在主页看到公开文章
    */
    async function bansomeone() {
        try {
            const response = await fetch(`/api/ban/${uid}`,{
               method: 'POST',
            }); 
            if (!response.ok) {
                throw new Error('封禁用户失败');
            }
            alert("封禁用户成功！")
            window.location.reload();
            
        } catch (err) {
           error = err.message;
        } finally {
            isLoading = false; 
        }
    }
    async function unbansomeone() {
        try {
            const response = await fetch(`/api/unban/${uid}`,{
               method: 'POST',
            }); 
            if (!response.ok) {
                throw new Error('解封用户失败');
            }
            alert("解封用户成功！")
            window.location.reload();
        } catch (err) {
           error = err.message;
        } finally {
            isLoading = false; 
        }
    }
    async function deletesomeone() {
      if(confirm('你确定要删除该用户吗？该用户的文章也会被一并删除')) {
        try {
            const response = await fetch(`/api/delete/${uid}`,{
               method: 'DELETE',
            }); 
            if (!response.ok) {
                throw new Error('删除用户失败');
            }
            alert("删除用户成功！")
            goto("/article/public")
        } catch (err) {
           error = err.message;
        } finally {
            isLoading = false; 
        }
      }
    }

</script>
<svelte:head>
  <title>{user.nickname}的主页</title>
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

    {#if isLoading} 
        <p>加载中...</p>
    {:else if error}
        <p style="color: red;">{error}</p>
    {:else}
      <div class = "user-con">
           <div class = "head">
            <h1>用户{user.nickname}的主页</h1>
           </div>
           <div class = "head">
               <p>用户的UID:{user.uid}<p></p>
               <p>用户的用户名:{user.username}<p></p>
               <p>创建时间：{user.createtime}</p>
               
              <div class = "view">
                {#if  !user.ban || user.ban && admin}
                   <a href={`/article/${user.uid}/public`} >
                   点我查看用户{user.nickname}公开的文章  </a>

                  <a href={`/article/${user.uid}/private`} >
                    点我查看用户{user.nickname}未公开的文章  </a>
  
                 {/if}
               </div>
               
              
          </div>

          <div class = "head">
            {#if currentUserUid === user.uid || admin}
          <!-- 如果当前用户是主页用户，则显示修改密码按钮 -->
          <button on:click={() => isEditingPassword = !isEditingPassword}>修改密码</button>
          {#if isEditingPassword}
              <!-- 显示修改密码表单 -->
              <div class="change-password-form">
                  <label for="new-password">新密码：</label>
                  <input type="password" id="new-password" bind:value={newPassword} placeholder="请输入新密码" required>
                  
                  <label for="confirm-password">确认密码：</label>
                  <input type="password" id="confirm-password" bind:value={confirmPassword} placeholder="请确认密码" required>

                  <button on:click={changePassword}>确认修改</button>
              </div>
             {/if}
          {/if}
          {#if user.ban}
            <h1>该用户已被封禁！！！</h1>
          {/if}

          {#if  admin && !user.admin}
          {#if !user.ban}
            <div class = "ban">
               <button on:click={bansomeone}>封禁该用户</button>
           </div>
           {:else}
             <div class = "ban">
               <button on:click={unbansomeone}>解封该用户</button>
          </div>
            {/if}
          {/if}

          {#if  admin && !user.admin}
            <div class = "ban">
               <button on:click={deletesomeone}>删除该用户</button>
           </div>
          {/if}
          


          </div>
          
          
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
    .head{
        text-align: center;

    }
    .search-container {
        display: flex;
    }

    .change-password-form {
		margin-top: 20px;
	}

	.change-password-form label {
		display: block;
		margin: 10px 0 5px;
	}

	.change-password-form input {
		width: 80%;
		padding: 8px;
		margin-bottom: 10px;
	}

	.change-password-form button {
		width: 80%;
		padding: 10px;
		background-color: #007bff;
		color: white;
		border: none;
		cursor: pointer;
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
.ban button {
  margin: 20px;
}
.head button {
  margin: 20px;
}

.view {
  display: flex;
  flex-direction: column;
}
</style>