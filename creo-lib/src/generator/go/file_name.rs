use crate::generator::core::{self, FileName};

pub struct FileNameGenerator;

impl core::FileNameGenerator for FileNameGenerator {
    fn generate_router_file_name(&self) -> FileName {
        FileName {
            path: "src/router",
            extension: "go",
        }
    }

    fn generate_service_call_file_name(&self) -> FileName {
        FileName {
            path: "src/service_calls",
            extension: "go",
        }
    }

    fn generate_main_file_name(&self) -> FileName {
        FileName {
            path: "src/main",
            extension: "go",
        }
    }
}
