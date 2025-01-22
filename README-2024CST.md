# tutorial

## 0 install git on windows

download from https://git-scm.com

```bat
git config --global credential.helper store
```

## 1 project directory 

create a working directory named as c:\kUser\t2024
```bat
set T2024=c:\kUser\t2024
md %T2024%
```

## 2 dev tool chains

### 2.1 golang

download from https://golang.google.cn

take tour from https://golang.google.cn/tour/welcome/1

```
# windows regedit
# Computer\HKEY_CURRENT_USER\Environment
# Computer\HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Control\Session Manager\Environment

#   key: GOPROXY, value: https://goproxy.cn,direct
#   key: GOSUMDB, value: sum.golang.google.cn

# linux .bashrc
# export GOPROXY=https://goproxy.cn,direct
# export GOSUMDB=sum.golang.google.cn


GOPROXY=https://goproxy.cn,direct
GOSUMDB=sum.golang.google.cn
```

### 2.2 vs code

download from https://code.visualstudio.com/

install plugins

```javascript
go: Go Team at Google
Svelte for VS Code: Svelte
```

### 2.3 nodejs
  install latest nodejs-LTS version from https://nodejs.org/en/download

### 2.4 pnpm

#### 2.4.1 install pnpm by
	
```bash
npm install -g pnpm

```

#### 2.4.2 using china source mirror for acceleration

```bat
pnpm config set registry https://registry.npmmirror.com/ 
```

## 3 svelte 

take tourial from https://svelte.dev/tutorial/svelte/welcome-to-svelte

### 3.1 init svelte

```bash
cd %T2024%
npx sv create fe
# Need to install the following packages:
# sv@0.6.10
# Ok to proceed? (y) 
#   y

# Which template would you like?
#   SvelteKit minimal (barebones scaffolding for your new app)

# Add type checking with Typescript?
#   Yes, using Javascript with JSDoc comments

# What would you like to add to your project? 
#   ◼ prettier
#   ◼ eslint
#   ◼ vitest
#   ◼ playwright
#   ◼ sveltekit-adapter

# What would you like to add to your project?
#   ● static (@sveltejs/adapter-static)

# Which package manager do you want to install dependencies with?
#   pnpm
```

### 3.2 update svelte packages

```bat
cd %T2024%\fe
pnpm update @sveltejs/kit @playwright/test @sveltejs/vite-plugin-svelte eslint-plugin-svelte prettier-plugin-svelte svelte svelte-check typescript vite

pnpm add  -D @sveltejs/adapter-static svelte-preprocess postcss sass bulma bulma-switch notyf material-icons material-symbols
```

### 3.3 alter %T2024%/fe/vite.config.js

```javascript
import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';
import path from 'path';

export default defineConfig({
	resolve: {
		alias: {
			'@src': path.resolve('./src')
		}
	},
	plugins: [sveltekit()],

	build: {
		sourcemap: true,
	},

	server: {
		host: 'localhost',
		port: 8080,
		open: true,
		proxy: {
			'/api': 'http://localhost:6616/',
		},
	}
});
```

### 3.4 alter  %T2024%/fe/svelte.config.js

```javascript

import adapter from '@sveltejs/adapter-static';

import preprocess from 'svelte-preprocess';



/** @type {import('@sveltejs/kit').Config} */
const config = {
	// Consult https://github.com/sveltejs/svelte-preprocess
	// for more information about preprocessors
	preprocess: preprocess(),

	onwarn: (warning, handler) => {
		if (warning.code.startsWith('a11y-')) {
			return;
		}
		handler(warning);
	},

	kit: {
		adapter: adapter({
			// default options are shown. On some platforms
			// these options are set automatically — see below
			pages: 'build',
			assets: 'build',
			fallback: 'index.html',
			precompress: true,
			strict: true,
		}),

	}
};

export default config;
```

### 3.5 alter  %T2024%/src/routes/+page.svelte

```html
<script>
	import { Notyf } from 'notyf';
	import { onMount } from 'svelte';

	let notyf;

	let name = $state('');
	let pwd = $state('');

	let login = (e) => {
		fetch(`/api/login?name=${name}&pwd=${pwd}`)
			.then((v) => {
				if (!v.ok) {
					console.log(v);
					return v;
				}

				return v.json();
			})
			.then((v) => {
				if (v.status !== 0) {
					console.log(v, decodeURIComponent(v.msg).replace(/\+/g, ' '));
					notyf.error('login failed');
					return;
				}

				notyf.success('login success');
			})
			.catch((e) => {
				console.log(e);
				if (e.message) {
					notyf.error(e.message);
				}
			});
	};

	onMount(() => {
		notyf = new Notyf({
			duration: 3000,
			className: 'x-notification',
			position: { x: 'right', y: 'top' }
		});
	});
</script>

<div class="main">
	<div class="title h1">login</div>
	<div>
		name: <input class="input" bind:value={name} type="text" />
	</div>
	<div>
		pwd: <input class="input" bind:value={pwd} type="password" />
	</div>
	<div>
		<button class="button is-primary" onclick={login}>login</button>
	</div>
</div>

<style lang="scss">
	@import 'notyf/notyf.min.css';
	@import 'bulma/css/bulma.css';

	.input {
		width: unset;
		max-width: unset;
	}

	.main {
		display: flex;
		flex-direction: column;
		align-items: center;

		div {
			padding: 1em;
		}
	}
</style>


```

## 4 init postgreSQL

download from https://www.postgresql.org/download/

documented by https://www.postgresql.org/docs/current/index.html

### using pgadmin as dbms tool

download from https://www.pgadmin.org/download/

documented by https://www.pgadmin.org/docs/pgadmin4/8.14/getting_started.html

### 4.1 create user/database/schema

```
user: testuser
database: testdb
schema: testuser

-- on database testuser run bellow script with user postgres
CREATE EXTENSION pgcrypto;
CREATE EXTENSION pg_stat_statements;
CREATE EXTENSION dblink;
```

```SQL
drop table if exists t_user;
create table t_user(id serial primary key,category varchar(32),account varchar(32),mobile_phone varchar(32), name varchar(32), user_token varchar(256),status varchar(4));
ALTER SEQUENCE t_user_id_seq RESTART WITH 20000;
insert into t_user(account,user_token,status) values
('tom',crypt('123',gen_salt('bf')),'00'),
('john',crypt('456',gen_salt('bf')),'00'),
('alic',crypt('789',gen_salt('bf')),'00');

select * from t_user;

drop table if exists t_article;
create table t_article(id serial primary key,category varchar(32),
  title varchar(32),author varchar(32), content varchar(32));
ALTER SEQUENCE t_article_id_seq RESTART WITH 20000;

```

### 4.2 alter pg_hba.conf 

add to last of file pg_hba.conf

```conf
host testdb testuser 0.0.0.0/0 md5
```

## 5 server

### 5.1 create directory

```bat
mkdir %T2024%\be
cd %T2024%\be
go mod init spider
```


### 5.2 main.go

```golang
package main

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v5"
)

var store = sessions.NewCookieStore([]byte("alongstory"))

func handler(w http.ResponseWriter, req *http.Request) {
	connUrl := "postgres://testuser:123456@localhost:6900/testdb"
	conn, err := pgx.Connect(context.Background(), connUrl)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Fprintf(w, `{"status":-1,"msg":"Unable to connect to database, %s"}`, url.QueryEscape(err.Error()))
		return
	}

	defer conn.Close(context.Background())

	var name, cert string
	n := req.URL.Query().Get("name")
	c := req.URL.Query().Get("pwd")

	session, _ := store.Get(req, "session-spider")

	if n == "" {
		n, _ = session.Values["name"].(string)
	}

	if c == "" {
		c, _ = session.Values["pwd"].(string)
	}

	if n != "" {
		session.Values["name"] = n
	}
	if c != "" {
		session.Values["pwd"] = c
	}

	session.Save(req, w)

	err = conn.QueryRow(context.Background(), "select account,user_token from t_user where account=$1 and user_token=crypt($2,user_token)", n, c).Scan(&name, &cert)
	if err != nil {
		fmt.Fprintf(w, `{"status":-1,"msg":"%s"}`, url.QueryEscape(err.Error()))
		return
	}

	fmt.Fprintf(w, `{"status":0,"msg":"%s"}`, "success")
}

func main() {

	http.HandleFunc("/api/login", handler)
	http.ListenAndServe(":6616", nil)

}

```

### 5.3 server build/run

```bat
cd %T2024%\be
go mod tidy
go run spider

```

## 6 that's all

good luck folks