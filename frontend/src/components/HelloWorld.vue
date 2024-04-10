<template>
  <div class="container">
    <h1>{{ msg }}</h1>
    <div class="cartoon-list">
      <h2>EchoSphere</h2>
      <ul>
        <li v-for="cartoon in cartoons" :key="cartoon.id">
          <div class="cartoon-item">
            <img :src="cartoon.image" alt="Cartoon Image">
            <div class="cartoon-details">
              <h3>{{ cartoon.name }}</h3>
              <p>{{ cartoon.description }}</p>
            </div>
          </div>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'HelloWorld',
  props: {
    msg: String
  },
  data() {
    return {
      cartoons: []
    };
  },
  mounted() {
    this.fetchCartoons();
  },
  methods: {
    fetchCartoons() {
      axios.get('http://127.0.0.1:8080/cartoon/')
          .then(response => {
            this.cartoons = response.data.cartoon;
          })
          .catch(error => {
            console.error('Error fetching cartoons:', error);
          });
    }
  }
}
</script>

<style scoped>
.container {
  max-width: 800px;
  margin: 0 auto;
  padding: 20px;
}

h1 {
  text-align: center;
  font-size: 24px;
  margin-bottom: 20px;
}

.cartoon-list {
  background-color: #f9f9f9;
  border-radius: 8px;
  padding: 20px;
}

h2 {
  font-size: 20px;
  margin-bottom: 10px;
}

ul {
  list-style-type: none;
  padding: 0;
}

.cartoon-item {
  display: flex;
  align-items: center;
  margin-bottom: 20px;
}

.cartoon-item img {
  width: 100px;
  height: auto;
  margin-right: 20px;
  border-radius: 8px;
}

.cartoon-details {
  flex: 1;
}

.cartoon-details h3 {
  margin: 0;
  font-size: 18px;
}

.cartoon-details p {
  margin: 5px 0;
  font-size: 14px;
  color: #666;
}
</style>
