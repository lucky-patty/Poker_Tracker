<template>
<v-container>
  <h1 class="display-3">Scoreboard</h1>
  <br>
  <v-divider></v-divider>
  <br>
  <br>
  <div class="text-center">
    <v-data-table item-key="name" class="elevation-1" loading loading-text="Loading... Please wait">
    </v-data-table>
  </div>
  <ScoreboardShow />
</v-container>
</template>

<style>
/* Enter and leave animations can use different */
/* durations and timing functions.              */
.slide-fade-enter-active {
  transition: all .3s ease;
}

.slide-fade-leave-active {
  transition: all .8s cubic-bezier(1.0, 0.5, 0.8, 1.0);
}

.slide-fade-enter,
.slide-fade-leave-to

/* .slide-fade-leave-active below version 2.1.8 */
  {
  transform: translateX(10px);
  opacity: 0;
}
</style>

<script>
import ScoreboardShow from '../components/ScoreboardTable.vue';
export default {
  components: {
    ScoreboardShow
  }
}
</script>
