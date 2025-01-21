// place files you want to import through the `$lib` alias in this folder.
import { goto } from '$app/navigation';

 export function goToPublicArticles() {
    goto('/article/public');
   }
   export function goTOMyArticles(currentUserUid) {
      goto(`/article/${currentUserUid}/myarticle`)
  }
  export function gotomypublicarticles(currentUserUid){
      goto(`/article/${currentUserUid}/public`)
  }
  export function gotomyprivatearticles(currentUserUid){
      goto(`/article/${currentUserUid}/private`)
  }
  export function gotocreatearticles(currentUserUid){
      goto(`/article/${currentUserUid}/create`)
  }
  export  function gotomypage(currentUserUid){
      goto(`/article/${currentUserUid}`)
  }