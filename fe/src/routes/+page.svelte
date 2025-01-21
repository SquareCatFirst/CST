<svelte:head>
  <title>登录</title>
</svelte:head>

<script>
	import { error } from "@sveltejs/kit";
    import { goto } from '$app/navigation';
    import { browser } from '$app/environment';
    let username = "";
    let password = "";
    let errorMessage = "";

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

    async function login() {
        if(!username || !password) {
            errorMessage = "用户名和密码不能为空！";
            return;
        }
        try {
            const response = await fetch("/api/login" , {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({ username, password }),
            });

            if(response.ok) {
                alert(`登录成功！`);
                goto('/article/public')
            } else {
                const data = await response.json();
                alert(data.message)
            }
        } catch (err) {
            alert('登录失败，请检查网络连接')
        }
    }

    async function gotoreg() {
        goto("/register")
    }
</script>

<main>
    <div class="login">
        <h1 class="head">登录</h1>
        <form class="list" on:submit|preventDefault={login}>
            <input type="text" placeholder="用户名" bind:value={username} />
            <input type="password" placeholder="密码" bind:value={password} />

            {#if errorMessage}
                <p class="error-message">{errorMessage}</p>
            {/if}

            <button type="submit">登录</button>
            <button type="button" on:click={gotoreg} class="register-btn">注册</button>
        </form>
    </div>
</main>

<style>
    .login {
        background-color: #fff;
        border-radius: 10px;
        box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
        width: 400px;
        padding: 30px;
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
        margin: 100px auto;
    }

    .head {
        font-size: 2rem;
        color: #333;
        margin-bottom: 20px;
    }

    .list {
        display: flex;
        flex-direction: column;
        width: 100%;
    }

    input[type="text"], input[type="password"] {
        padding: 10px;
        margin: 10px 0;
        font-size: 1rem;
        border: 1px solid #ccc;
        border-radius: 5px;
        background-color: #fafafa;
        transition: border 0.3s ease;
    }

    input[type="text"]:focus, input[type="password"]:focus {
        border-color: #4CAF50;
        outline: none;
    }

    .error-message {
        color: red;
        font-size: 0.9rem;
        margin-bottom: 10px;
    }

    button {
        padding: 10px;
        font-size: 1rem;
        border: none;
        border-radius: 5px;
        background-color: #4CAF50;
        color: white;
        margin: 10px 0;
        cursor: pointer;
        transition: background-color 0.3s ease;
    }

    button:hover {
        background-color: #45a049;
    }

    .register-btn {
        background-color: #ccc;
    }

    .register-btn:hover {
        background-color: #bbb;
    }
</style>
