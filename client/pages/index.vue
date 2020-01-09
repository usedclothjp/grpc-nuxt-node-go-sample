<template>
  <div class="hello">
    <input type="text" placeholder="num1" v-model="num1">
    <input type="text" placeholder="num2" v-model="num2">
    <button @click="submitClick">submit</button>
    <p>{{ this.result }}</p>
  </div>
</template>

<script>
import Axios from "axios";
export default {
  data() {
    return {
      num1: null,
      num2: null,
      result: 0
    };
  },
  methods: {
    async submitClick() {
      const params = {num1: this.num1, num2: this.num2};
      const { data } = await Axios.get('api/rpc', {params})
      this.result = data.num
      console.log(this.result)
    },
    dosomething(e) {
      console.log('asdfewafdewafa')
    },
    addCount (e) {
      console.log('=================')
      // this.$store.commit('counter/add')
    },
  }
};
</script>
