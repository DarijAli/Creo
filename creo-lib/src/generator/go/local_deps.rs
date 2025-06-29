pub fn get_local_handler_dependencies(
    lib_dir: impl AsRef<std::path::Path>,
) -> std::io::Result<Vec<String>> {
    let mut deps = Vec::default();
    for entry in lib_dir.as_ref().read_dir()? {
        let entry = entry?;
        let ft = entry.file_type()?;
        if ft.is_dir() {
            let dep_name = entry.file_name();
            let dep_str_name = dep_name
                .to_str()
                .expect("directory name should be valid UTF-8");

            deps.push(format!("require templates/go/lib/{} v0.0.0", dep_str_name));
            deps.push(format!(
                "replace templates/go/lib/{} v0.0.0 => ./lib/{}",
                dep_str_name, dep_str_name
            ));
        } else {
            log::debug!("Skipping entry {}", entry.path().display());
        }
    }

    Ok(deps)
}
