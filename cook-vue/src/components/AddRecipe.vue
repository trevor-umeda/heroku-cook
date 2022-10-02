<template>
  <div class="submit-form">
    <div v-if="!submitted">
      <div class="form-group">
        <label for="title">Title</label>
        <input
            type="text"
            class="form-control"
            id="title"
            required
            v-model="recipe.title"
            name="title"
        />
      </div>

      <div class="form-group">
        <label for="link">Link</label>
        <input
            class="form-control"
            id="link"
            required
            v-model="recipe.link"
            name="link"
        />
      </div>

      <div class="form-group">
        <label for="tags">Tags</label>
        <input
            class="form-control"
            id="tags"
            required
            v-model="recipe.tags"
            name="tags"
        />
      </div>

      <div class="form-group">
        <label for="rating">Rating</label>
        <input
            class="form-control"
            id="rating"
            required
            v-model="recipe.rating"
            name="rating"
        />
      </div>

      <button @click="saveRecipe" class="btn btn-success">Submit</button>
    </div>

    <div v-else>
      <h4>You submitted successfully!</h4>
      <button class="btn btn-success" @click="newRecipe">Add</button>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import RecipeDataService from "@/services/RecipeDataService";
import type Recipe from "@/types/Recipe";
import type ResponseData from "@/types/ResponseData";
export default defineComponent({
  name: "add-recipe",
  data() {
    return {
      recipe: {
        id: 0,
        title: "",
        link: "",
        tags: "",
        rating: 0
      },
      submitted: false
    };
  },
  methods: {
    saveRecipe() {
      var data = {
        title: this.recipe.title,
        link: this.recipe.link,
        tags: this.recipe.tags,
        rating: this.recipe.rating
      };

      RecipeDataService.create(data)
          .then(response => {
            this.recipe.id = response.data.id;
            console.log(response.data);
            this.submitted = true;
          })
          .catch(e => {
            console.log(e);
          });
    },

    newRecipe() {
      this.submitted = false;
      this.recipe = {} as Recipe;
    }
  }
});
</script>

<style>
.submit-form {
  max-width: 300px;
  margin: auto;
}
</style>