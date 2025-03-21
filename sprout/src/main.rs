use axum::{
    routing::{get, post},
    Router,
    response::Html,
    extract::Extension,
    Json,
};
use async_graphql::{
    http::{playground_source, GraphQLPlaygroundConfig},
    Schema, EmptySubscription, Object, Context, Result,
};
use async_graphql_axum::{GraphQLRequest, GraphQLResponse};
use std::sync::Arc;

// Definicja struktury Query
struct QueryRoot;

#[Object]
impl QueryRoot {
    async fn hello(&self) -> &'static str {
        "Hello, world!"
    }
}

// Definicja struktury Mutation
struct MutationRoot;

#[Object]
impl MutationRoot {
    async fn set_name(&self, ctx: &Context<'_>, name: String) -> Result<String> {
        let mut data = ctx.data_unchecked::<AppState>();
        data.name = name.clone();
        Ok(name)
    }
}

// Stan aplikacji
struct AppState {
    name: String,
}

// Typ schematu GraphQL
type SchemaType = Schema<QueryRoot, MutationRoot, EmptySubscription>;

#[tokio::main]
async fn main() {
    // Inicjalizacja schematu GraphQL
    let schema = Schema::build(QueryRoot, MutationRoot, EmptySubscription)
        .data(AppState {
            name: "World".to_string(),
        })
        .finish();

    // Tworzenie routera Axum
    let app = Router::new()
        .route("/graphql", post(graphql_handler))
        .route("/playground", get(graphql_playground))
        .layer(Extension(Arc::new(schema)));

    println!("Playground: http://localhost:8000/playground");

    // Uruchomienie serwera
    axum::Server::bind(&"0.0.0.0:8000".parse().unwrap())
        .serve(app.into_make_service())
        .await
        .unwrap();
}

// Obsługa zapytań GraphQL
async fn graphql_handler(
    Extension(schema): Extension<Arc<SchemaType>>,
    req: Json<GraphQLRequest>,
) -> Json<GraphQLResponse> {
    Json(schema.execute(req.0).await.into())
}

// Obsługa GraphQL Playground
async fn graphql_playground() -> Html<String> {
    Html(playground_source(GraphQLPlaygroundConfig::new("/graphql")))
}