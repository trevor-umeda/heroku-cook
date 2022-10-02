import http from "../http-common";

class RecipeDataService {
    getAll(): Promise<any> {
        return http.get("/recipes");
    }

    get(id: any): Promise<any> {
        return http.get(`/recipes/${id}`);
    }

    create(data: any): Promise<any> {
        return http.post("/recipes", data);
    }

    update(id: any, data: any): Promise<any> {
        return http.put(`/recipes/${id}`, data);
    }

    findByTitle(title: string): Promise<any> {
        return http.get(`/recipes?title=${title}`);
    }
}

export default new RecipeDataService();