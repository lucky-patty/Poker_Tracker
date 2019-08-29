<template>
<v-app>
  <v-navigation-drawer clipped v-model="drawer" app dark>
    <v-list dense nav>
      <v-list-item v-for="item in items" :key="item.title" @click="" :to="item.route" link>
        <v-list-item-icon>
          <v-icon>{{item.icon}}</v-icon>
        </v-list-item-icon>
        <v-list-item-content>
          <v-list-item-title>{{item.title}}</v-list-item-title>
        </v-list-item-content>
      </v-list-item>
    </v-list>
  </v-navigation-drawer>
  <v-app-bar app fixed clipped-left dark>
    <v-app-bar-nav-icon @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
    <v-toolbar-title class="white--text">Poker Score Tracker</v-toolbar-title>
    <v-spacer></v-spacer>
    <v-toolbar-items>
      <v-btn icon>
        <v-icon>refresh</v-icon>
      </v-btn>
    </v-toolbar-items>
  </v-app-bar>
  <v-content>
    <transition name="router-anim" enter-active-class="animated fadeInUp" leave-active-class="animated fadeOutDown">
      <router-view />
    </transition>
  </v-content>
</v-app>
</template>

<style>
@import "https://cdnjs.cloudflare.com/ajax/libs/animate.css/3.7.2/animate.min.css";
</style>

<script>
export default {
  data: () => ({
    drawer: false,
    items: [{
        title: 'Home',
        icon: 'home',
        route: '/',
      },
      {
        title: 'Tracking',
        icon: 'games',
        route: '/track',
      },
      {
        title: 'Scoreboard',
        icon: 'dashboard',
        route: '/scoreboard',
      },
      {
        title: 'Setting',
        icon: 'settings',
        route: '/about',
      },
    ]
  }),
  props: {
    source: String
  }
};
</script>
