<template>
<div>
  <v-navigation-drawer clipped v-model="drawer" app dark>
    <v-list dense nav>
      <v-list-item v-for="item in items" :key="item.title" link>
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
      <v-btn icon flat>
        <v-icon>apps</v-icon>
      </v-btn>
      <v-btn icon>
        <v-icon>refresh</v-icon>
      </v-btn>
    </v-toolbar-items>
  </v-app-bar>

</div>
</template>

<script>
export default {
  data: () => ({
    drawer: false,
    items: [{
      title: 'Scoreboard',
      icon: 'dashboard',
    },
    {
      title: 'Setting',
      icon: 'settings',
    },
    ],
  }),
  props: {
    source: String,
  },
};
</script>
