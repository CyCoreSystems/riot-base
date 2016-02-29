require('./nav.tag')
require('./footer.tag')

require('./pages/home.tag')

<app class="page">
   <nav></nav>
   <div id='main'></div>
   <footer></footer>

   riot.route('/',() => riot.mount('#main','home'))

</app>
