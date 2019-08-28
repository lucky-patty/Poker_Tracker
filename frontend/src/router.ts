import Vue from 'vue';
import Router from 'vue-router';
import Home from './views/Home.vue';
import Scoreboard from './views/Scoreboard.vue';
import PokerIndex from './views/PokerIndex.vue';

Vue.use(Router);

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/home',
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
    {
      path: '/scoreboard',
      name: 'scoreboard',
      component: Scoreboard,
    },
    {
      path: '/poker',
      name: 'poker index',
      component: PokerIndex,
    },
  ],
});
