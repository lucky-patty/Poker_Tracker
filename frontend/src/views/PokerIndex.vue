<template>
<v-container>
  <h1 class="display-3">Welcome Back</h1>
  <v-layout text-center>
    <v-card flat>
      <v-img class="white--text" height="300px" src="https://picsum.photos/id/1010/5184/3456">
        <v-card-title class="align-end fill-height">Welcome to Proker Tracker</v-card-title>
      </v-img>
      <v-card-text>
        <span>
          This application was made because I find it hard to track poker score!
        </span>
      </v-card-text>
      <v-card-actions>
        <router-link to="/about">
          <v-btn text>
            Start Playing!
          </v-btn>
        </router-link>
      </v-card-actions>
    </v-card>
  </v-layout>
  <h1 class="display-1">Activity</h1>
  <v-layout>
    <v-card flat>
      <v-sparkline :value="value" :gradient="gradient" :smooth="radius || false" :padding="padding" :line-width="lineWidth" :stroke-linecap="lineCap" :gradient-direction="gradientDirection" :fill="fill" :type="type" :auto-line-width="autoLineWidth"
        auto-drawer :show-labels="showLabels" :label-size="labelSize">
      </v-sparkline>
      <v-divider></v-divider>
    </v-card>
  </v-layout>

</v-container>
</template>

<script>
const gradients = [
  ['#222'],
  ['#42b3f4'],
  ['red', 'orange', 'yellow'],
  ['purple', 'violet'],
  ['#00c6ff', '#F0F', '#FF0'],
  ['#f72047', '#ffd200', '#1feaea'],
];
export default {
  data: () => ({
    showLabels: false,
    lineWidth: 2,
    labelSize: 7,
    radius: 10,
    padding: 8,
    lineCap: 'round',
    gradient: gradients[5],
    value: [0, 2, 5, 9, 5, 10, 3, 5, -4, -10, 1, 8, 2, 9, 0],
    gradientDirection: 'top',
    gradients,
    fill: false,
    type: 'trend',
    autoLineWidth: false,
  }),
};
</script>
