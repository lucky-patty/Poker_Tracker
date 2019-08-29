<template>
<v-card flat>
  <v-navigation-drawer clipped fixed v-model="drawer" dark>
    <v-list dense>
      <v-list-tile>
        <v-list-tile-action>
          <v-icon>dashboard</v-icon>
        </v-list-tile-action>
        <v-list-tile-content>
          <v-list-tile-title>scoreboard</v-list-tile-title>
        </v-list-tile-content>
      </v-list-tile>
    </v-list>
  </v-navigation-drawer>
  <v-layout row wrap>
    <v-toolbar fixed clipped-left dark>
      <v-toolbar-side-icon @click.stop="drawer = !drawer">
      </v-toolbar-side-icon>
      <v-toolbar-title>Poker Score Tracker</v-toolbar-title>
      <v-spacer></v-spacer>
    </v-toolbar>
  </v-layout>
  <!-- </v-container> -->
</v-card>
</template>

<script>
export default {
  data() {
    return {
      drawer: true,
      items: [{
        title: 'Scoreboard',
        icon: 'dashboard',
      },
      {
        title: 'Setting',
        icon: 'setting',
      },
      ],
      right: null,
    };
  },
};
</script>
