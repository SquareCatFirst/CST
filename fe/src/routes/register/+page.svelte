<svelte:head>
  <title>注册</title>
</svelte:head>
<script>
	import { goto } from "$app/navigation";
  import { browser } from '$app/environment';
    let nickname = '';
    let username = '';  
    let phone = '';
    let email = '';  
    let password = '';
    let confirmPassword = '';
    let isadmin = false;


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
      if(currentUserUid !== '' && currentUserUid !== null) {
        goto('/article/public')
      }
    }
    if(browser) ifcookie()
    
    let errors = {
      nickname: '',
      username: '',  
      phone: '',
      email: '', 
      password: '',
      confirmPassword: '',
    };
  
    const validateForm = () => {
      let isValid = true;
  
      // 清空错误信息
      errors = {
        nickname: '',
        username: '',  
        phone: '',
        email: '',  
        password: '',
        confirmPassword: '',
      };
      // 验证昵称不能存在空格
      if(nickname.includes(' ')) {
        errors.nickname = '昵称不能存在空格';
        isValid = false;
      }
      // 验证用户名不能有空格
      if(username.includes(' ')) {
        errors.username = '用户名不能存在空格';
        isValid = false;
      }

      // 验证昵称
      if (!nickname) {
        errors.nickname = '昵称不能为空';
        isValid = false;
      }
  
      // 验证用户名（不能为空）
      if (!username) {
        errors.username = '用户名不能为空';
        isValid = false;
      }
  
      // 验证手机号（简单的正则验证）
      const phoneRegex = /^1[3-9]\d{9}$/;
      if (!phone || !phoneRegex.test(phone)) {
        errors.phone = '请输入有效的手机号';
        isValid = false;
      }
  
      // 验证邮箱（简单的正则验证）
      const emailRegex = /^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,6}$/;
      if (!email || !emailRegex.test(email)) {
        errors.email = '请输入有效的邮箱地址';
        isValid = false;
      }
  
      // 验证密码（至少6个字符）
      if (!password || password.length < 6) {
        errors.password = '密码必须至少6个字符';
        isValid = false;
      }
  
      // 验证确认密码
      if (password !== confirmPassword) {
        errors.confirmPassword = '密码与确认密码不一致';
        isValid = false;
      }
  
      return isValid;
    };
  
    const handleSubmit = async () => {
      if (validateForm()) {
        // 提交表单
        try {
          const response = await fetch('/api/register', {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json',
            },
            body: JSON.stringify({
              nickname,
              username,  // 发送用户名
              phone,
              email,  // 发送邮箱
              password,
              isadmin,
            }),
          });

          const data = await response.json();
          
          if (response.ok) {
            alert('注册成功')
            goto("/")
          } else {
            alert(data.message)
          }
        } catch (error) {
          alert(error)
        }
      } else {
        console.log('表单有错误，无法提交');
      }
    };

    async function goBack() {
      goto("/")
    }
  </script>
  
  <main>
    <div class="form-container">

      <button class="back-btn" on:click={goBack}>返回</button>
      <h2>注册</h2>
        
    <!-- 用户名 -->
    <div class="form-group">
      <label for="username">用户名</label>
      <input type="text" id="username" bind:value={username} placeholder="请输入用户名" />
     {#if errors.username}
          <div class="error">{errors.username}</div>
        {/if}
    </div>

      <!-- 昵称 -->
      <div class="form-group">
        <label for="nickname">昵称</label>
        <input type="text" id="nickname" bind:value={nickname} placeholder="请输入昵称" />
        {#if errors.nickname}
          <div class="error">{errors.nickname}</div>
        {/if}
      </div>
  
      <!-- 手机 -->
      <div class="form-group">
        <label for="phone">手机号</label>
        <input type="text" id="phone" bind:value={phone} placeholder="请输入手机号" />
        {#if errors.phone}
          <div class="error">{errors.phone}</div>
        {/if}
      </div>
  
      <!-- 邮箱 -->
      <div class="form-group">
        <label for="email">邮箱</label>
        <input type="email" id="email" bind:value={email} placeholder="请输入邮箱" />
        {#if errors.email}
          <div class="error">{errors.email}</div>
        {/if}
      </div>
  
      <!-- 密码 -->
      <div class="form-group">
        <label for="password">密码</label>
        <input type="password" id="password" bind:value={password} placeholder="请输入密码" />
        {#if errors.password}
          <div class="error">{errors.password}</div>
        {/if}
      </div>
  
      <!-- 确认密码 -->
      <div class="form-group">
        <label for="confirmPassword">确认密码</label>
        <input type="password" id="confirmPassword" bind:value={confirmPassword} placeholder="确认密码" />
        {#if errors.confirmPassword}
          <div class="error">{errors.confirmPassword}</div>
        {/if}
      </div>
  
      <!-- 提交按钮 -->
      <button on:click={handleSubmit}>注册</button>
    </div>
  </main>
  
  <style>
    .form-container {
      width: 800px;
      display: flex;
      flex-direction: column;
      border: 2px solid rgb(114, 97, 97);
      border-radius: 10px;
      align-items: center;
      margin: 100px auto;
    }
  
    .form-group {
      margin-bottom: 20px;
      display: flex;
      flex-direction: column;
    }
  
    input[type="text"],
    input[type="password"],
    input[type="email"] {
      width: 100%;
      padding: 10px;
      font-size: 16px;
      border-radius: 5px;
      border: 1px solid #ccc;
      margin-right: 20px;
      margin-left: 20px;
    }
  
    .error {
      color: red;
      font-size: 12px;
      margin-top: 5px;
    }
  
    button {
      width: 100%;
      padding: 10px;
      font-size: 16px;
      cursor: pointer;
      border-radius: 5px;
      border: none;
      background-color: #4CAF50;
      color: white;
      transition: background-color 0.3s;
    }
  
    button:hover {
      background-color: #45a049;
    }

    .back-btn {
    position: absolute;
    top: 10px;
    left: 10px;
    padding: 5px 10px;
    background-color: #ccc;
    border: none;
    border-radius: 5px;
    cursor: pointer;
  }
  </style>
  