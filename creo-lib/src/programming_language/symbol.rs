use super::ProgrammingLanguage::{self, *};
use crate::generator::{self, core::SymbolGenerator};

impl ProgrammingLanguage {
    pub(crate) fn as_symbol_generator(&self) -> &dyn SymbolGenerator {
        match self {
            Python(_) => &generator::python::SymbolGenerator,
            Rust(_) => &generator::rust::SymbolGenerator,
            Go(_) => &generator::go::SymbolGenerator,
        }
    }
}
