use rand_derive::Rand;

use crate::{generator::core::FrameworkGenerator, template};

pub mod gin;

#[derive(Rand)]
pub enum Frameworks {
    Gin,
}
use Frameworks::*;

impl FrameworkGenerator for Frameworks {
    fn to_faker(&self) -> &dyn template::Fakeable {
        match self {
            Gin => &gin::Faker,
        }
    }

    fn to_router_generator(&self) -> &dyn template::RouterGenerator {
        match self {
            Gin => &gin::RouterGenerator,
        }
    }

    fn to_service_calls_generator(&self) -> &dyn template::ServiceCallGenerator {
        match self {
            Gin => &gin::ServiceCallGenerator,
        }
    }

    fn to_main_generator(&self) -> &dyn template::MainGenerator {
        match self {
            Gin => &gin::MainGenerator,
        }
    }

    fn get_framework_requirements(&self) -> Vec<&'static str> {
        match self {
            Gin => gin::get_framework_dependencies(),
        }
    }

    fn get_docker_entrypoint(&self) -> &'static str {
        match self {
            Gin => gin::DOCKER_ENTRYPOINT,
        }
    }
}
