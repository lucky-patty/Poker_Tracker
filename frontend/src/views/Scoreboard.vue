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

<script>
import ScoreboardShow from '../components/ScoreboardTable.vue';
export default {
  components: {
    ScoreboardShow
  }
}
</script>
