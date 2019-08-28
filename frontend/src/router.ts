import Vue from 'vue';
import Router from 'vue-router';
import Home from './views/Home.vue';

// Our Pages
// import PokerIndex from './views/PokerIndex.vue';
import PokerScore from './views/ShowScore.vue';

Vue.use(Router);

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home,
    },
    {
      path: '/about',
      name: 'about',
      // route level code-splitting
      // this generates a separate chunk (about.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import(/* webpackChunkName: "about" */ './views/About.vue'),
    },
    // {
    //   path: '/poker',
    //   name: 'poker index',
    //   component: PokerIndex,
    // },
    {
      path: '/score',
      name: 'poker score',
      component: PokerScore,
    },
  ],
});
