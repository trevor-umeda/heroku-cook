<template>
  <div class="recipes">
    <div class="recipe" v-for="recipe in recipes" :key="recipe.id">
      <a :href="recipe.link"><h2>{{ recipe.title }}  ({{recipe.rating}})</h2></a>
      <p>{{ recipe.link }}</p>
    </div>
  </div>
</template>
<script lang="ts">
import { defineComponent } from "vue";
import RecipeDataService from "@/services/RecipeDataService";
import type Recipe from "@/types/Recipe";
import type ResponseData from "@/types/ResponseData";

export default defineComponent({
  name: "recipes-list",
  data() {
    return {
      recipes: [] as Recipe[],
      id: "",
      link: "",
      title: "",
      rating: ""
    }
  },
  methods: {
    retrieveRecipes() {
      RecipeDataService.getAll()
          .then((response: ResponseData) => {
            this.recipes = response.data;
            console.log(response.data);
          })
          .catch((e: Error) => {
            console.log(e);
          });
    }
  },
  mounted() {
    this.retrieveRecipes();

  }
})
</script>
<style>
@media (min-width: 1024px) {
  .about {
    min-height: 100vh;
    display: flex;
    align-items: center;
  }
}
</style>
