mod data_type;
mod file_name;
mod frameworks;
mod local_deps;
mod symbol;

pub use data_type::DataTypeMapper;
pub use file_name::FileNameGenerator;
pub use frameworks::Frameworks;
pub use local_deps::get_local_handler_dependencies;
pub use symbol::SymbolGenerator;

pub const DOCKERFILE_TEMPLATE_PATH: &str = "go/Dockerfile.mgt";
pub const DEPENDENCY_FILE_NAME: &str = "go.mod";
pub const DEPENDENCY_FILE_TEMPLATE_PATH: &str = "go/go.mgt";
