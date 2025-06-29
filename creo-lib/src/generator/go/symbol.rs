use crate::generator::core;

pub struct SymbolGenerator;

impl core::SymbolGenerator for SymbolGenerator {
    fn generate_array_item_function_name(&self, name: &str) -> String {
        format!("{}_Item", name)
    }

    fn generate_object_property_function_name(&self, name: &str, prop_name: &str) -> String {
        format!("{}_Prop_{}", name, prop_name)
    }

    fn generate_service_calls_function_name(
        &self,
        endpoint: crate::graph::EndpointIndex,
    ) -> String {
        format!("ServiceCallsEndpoint{}", endpoint.0)
    }

    fn generate_service_call_function_import(
        &self,
        file_path: &str,
        _function_name: &str,
    ) -> String {
        let package = file_path
            .rsplit_once('/')
            .map(|(_, p)| p)
            .unwrap_or(file_path);
        format!("./lib/{}", package)
    }

    fn generate_individual_service_call_function_name(
        &self,
        call: crate::application::ServiceCallEdge,
    ) -> String {
        format!(
            "ServiceCallEndpoint{}_ToEndpoint{}",
            call.source.0, call.target.0
        )
    }

    fn generate_operation_function_name(&self, endpoint: crate::graph::EndpointIndex) -> String {
        format!("OperationEndpoint{}", endpoint.0)
    }

    fn generate_handler_function_import(&self, import_path: &str, _function_name: &str) -> String {
        format!("templates/go/lib/{}/src", import_path)
    }

    fn generate_query_data_function_name(
        &self,
        service_call: crate::application::ServiceCallEdge,
    ) -> String {
        format!(
            "QueryDataCallEndpoint{}_ToEndpoint{}",
            service_call.source.0, service_call.target.0
        )
    }

    fn generate_parameter_function_name(
        &self,
        service_call: crate::application::ServiceCallEdge,
        param_name: &str,
    ) -> String {
        format!(
            "CallEndpoint{}_ToEndpoint{}_Param{}",
            service_call.source.0, service_call.target.0, param_name
        )
    }
}
