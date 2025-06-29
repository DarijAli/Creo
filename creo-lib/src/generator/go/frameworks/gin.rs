use crate::template;

pub const DOCKER_ENTRYPOINT: &str =
    r#"[ "gin", "run", "main.go", "--host", "0.0.0.0", "--port", "80" ]"#;

pub struct Faker;

impl template::Fakeable for Faker {
    fn get_string_fake(
        &self,
        _string_validation: &openapiv3::StringType,
    ) -> template::FakeFunction {
        // direct faker.Name() call, no args needed
        template::FakeFunction::new("faker.Name()".into(), String::new())
    }

    fn get_number_fake(
        &self,
        _number_validation: &openapiv3::NumberType,
    ) -> template::FakeFunction {
        template::FakeFunction::new("faker.Float64()".into(), String::new())
    }

    fn get_integer_fake(
        &self,
        _integer_validation: &openapiv3::IntegerType,
    ) -> template::FakeFunction {
        template::FakeFunction::new("faker.Int()".into(), String::new())
    }

    fn get_object_fake(&self, function_name: &str) -> template::FakeFunction {
        template::FakeFunction::new(function_name.into(), String::new())
    }

    fn get_array_fake(&self, function_name: &str) -> template::FakeFunction {
        template::FakeFunction::new(function_name.into(), String::new())
    }

    fn get_boolean_fake(
        &self,
        _boolean_validation: &openapiv3::BooleanType,
    ) -> template::FakeFunction {
        template::FakeFunction::new("faker.Bool()".into(), String::new())
    }
}

pub struct RouterGenerator;

impl template::RouterGenerator for RouterGenerator {
    fn create_router_template(&self) -> template::RouterTemplate {
        template::RouterTemplate {
            template_dir: "go/gin/router",
            root_template_name: "router",
        }
    }
}

pub struct ServiceCallGenerator;

impl template::ServiceCallGenerator for ServiceCallGenerator {
    fn create_service_call_template(&self) -> template::ServiceCallTemplate {
        template::ServiceCallTemplate {
            template_dir: "go/gin/service_calls",
            root_template_name: "service_calls",
        }
    }
}

pub struct MainGenerator;

impl template::MainGenerator for MainGenerator {
    fn create_main_template(&self) -> template::MainTemplate {
        template::MainTemplate {
            template_dir: "go/gin",
            root_template_name: "main",
            auxiliry_template_names: &[template::AuxiliryTemplate {
                template_name: "http_client",
                file_name: "src/http_client.go",
            }],
        }
    }
}

pub fn get_framework_dependencies() -> Vec<&'static str> {
    vec![
        "require github.com/gin-gonic/gin v1.9.1",
        "require github.com/go-faker/faker/v4 v4.6.1",
    ]
}
