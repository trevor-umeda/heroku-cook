<script setup lang="ts">
import { RouterLink, RouterView } from "vue-router";
import Landing from "./components/Landing.vue";
import AddRecipe from "@/components/AddRecipe.vue";
</script>

<template>
  <header>
    <div class="wrapper">
      <Landing msg="Cooking App" />
      <AddRecipe />
      <nav>
        <RouterLink to="/">Home</RouterLink>
        <RouterLink to="/list">Recipes</RouterLink>
      </nav>
      <main>
        <div class="row">
          <div class="col-8">
            <h3>Draggable table</h3>

            <table class="table table-striped">
              <thead class="thead-dark">
              <tr>
                <th scope="col">Id</th>
                <th scope="col">Name</th>
                <th scope="col">Sport</th>
              </tr>
              </thead>
              <draggable v-model="list" tag="tbody" item-key="name">
                <template #item="{ element }">
                  <tr>
                    <td scope="row">{{ element.id }}</td>
                    <td>{{ element.name }}</td>
                    <td>{{ element.sport }}</td>
                  </tr>
                </template>
              </draggable>
            </table>
          </div>

          <rawDisplayer class="col-3" :value="list" title="List" />
        </div>
      </main>
    </div>
  </header>

  <RouterView />
</template>
<script lang="ts">
import axios from "axios";
import { defineComponent } from "vue";

export default defineComponent({
  name: "App",
  data() {
    return {
      recipeName: "",
      recipeLink: "",
      recipeTags: "",
      recipeRating: 0
    };
  },
  methods: {
    async addRecipe() {
      console.log("Nice");
      const res = await axios.post(`/api/recipes`, {
        title: this.recipeName,
        link: this.recipeLink,
        tags: this.recipeTags,
        rating: this.recipeRating,
      });
    },
  },
});
</script>
<style scoped>
.draggable-list {
  background: #3f51b5;
  color: #fff;
  border: 1px solid;
  height: 50vh;
}
.list-item {
  margin: 10px;
  padding: 40px;
  cursor: pointer;
  font-size: 18px;
  border-radius: 5px;
  background: #f44336;
  display: inline-block;
}

header {
  line-height: 1.5;
  max-height: 100vh;
}

.logo {
  display: block;
  margin: 0 auto 2rem;
}

nav {
  width: 100%;
  font-size: 12px;
  text-align: center;
  margin-top: 2rem;
}

nav a.router-link-exact-active {
  color: var(--color-text);
}

nav a.router-link-exact-active:hover {
  background-color: transparent;
}

nav a {
  display: inline-block;
  padding: 0 1rem;
  border-left: 1px solid var(--color-border);
}

nav a:first-of-type {
  border: 0;
}

.recipeInput {
  display:block;

}

@media (min-width: 1024px) {
  header {
    display: flex;
    place-items: center;
    padding-right: calc(var(--section-gap) / 2);
  }

  .logo {
    margin: 0 2rem 0 0;
  }

  header .wrapper {
    place-items: flex-start;
    flex-wrap: wrap;
  }

  nav {
    text-align: left;
    margin-left: -1rem;
    font-size: 1rem;

    padding: 1rem 0;
    margin-top: 1rem;
  }
}
</style>
